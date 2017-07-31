package pm

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/nsq"
	"golang-learn/examples/64_manager/sqlcon"
	"golang-learn/examples/64_manager/web"
	"os"
	"os/signal"
	"syscall"
)

// IManager ... manager interface
type IManager interface {
	// Processes
	GetProcess(key string) IProcess
	AddProcess(key string, process IProcess) error
	RemProcess(key string) (IProcess, error)

	// Configurations
	GetConfig(key string) IConfig
	AddConfig(key string, config IConfig) error
	RemConfig(key string) (IConfig, error)

	// Connections
	GetConnection(key string) (*sql.DB, error)
	AddConnection(key string, controller *sqlcon.SQLConController) error
	RemConnection(key string) (*sql.DB, error)

	// Web Servers
	AddWebServer(key string, webServer web.IWebController)
	AddWebServerRoute(key string, method string, route string, handler func(context echo.Context) error)
	StartWebServer(key string) error
	StopWebServer(key string) error

	// NEW Instances
	NewSQLConnection(config *sqlcon.Config) (*sqlcon.SQLConController, error)
	NewNSQConsumer(config *nsq.Config, handler nsq.IHandler) (nsq.IConsumer, error)
	NewNSQProducer(config *nsq.Config) (nsq.IProducer, error)
	NewWEBServer(config *web.Config) (web.IWebController, error)
	NewSimpleConfig(path string, file string, extension string) (IConfig, error)

	// MANAGER
	Start() error
	Stop() error
}

// manager ... manager structure
type manager struct {
	processController map[string]*ProcessController
	configController  map[string]*ConfigController
	sqlConController  map[string]*sqlcon.SQLConController
	httpController    map[string]web.IWebController

	control chan int
	started bool
}

// NewManager ... create a new manager
func NewManager() (IManager, error) {

	return &manager{
		processController: make(map[string]*ProcessController),
		configController:  make(map[string]*ConfigController),
		sqlConController:  make(map[string]*sqlcon.SQLConController),
		httpController:    make(map[string]web.IWebController),

		control: make(chan int),
	}, nil
}

// Start ... starts and blocks until it receives a signal in its control channel or a SIGTERM,
func (instance *manager) Start() error {
	log.Infof("manager, starting")
	instance.started = true

	// listen for termination signals
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	// launch every process in a separeted process
	for name, process := range instance.processController {
		log.Infof("manager, starting process [process:%s]", name)

		go instance.launch(name, process)

		log.Infof("manager, started process [process:%s]", name)
	}

	select {
	case <-termChan:
		log.Infof("manager, received term signal")
	case <-instance.control:
		log.Infof("manager, received shutdown signal")
	}

	instance.Stop()

	return nil
}

// Stop ... stop all processes and stops the manager
func (instance *manager) Stop() error {
	if instance.started {
		log.Infof("manager, stopping")

		for key, controller := range instance.processController {
			log.Infof("manager, stopping process [process:%s]", key)
			if err := controller.process.Stop(); err != nil {
				log.Error(err, fmt.Sprintf("error stopping process [process:%s]", key))
			}
			log.Infof("manager, waiting for process to terminate [process:%s]", key)
			<-controller.control
			close(controller.control)
			delete(instance.processController, key)
			log.Infof("manager, stopped process [process:%s]", key)
		}

		instance.started = false
		log.Infof("manager, stopped")
	}

	return nil
}
