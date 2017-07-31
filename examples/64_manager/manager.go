package pm

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/config"
	"golang-learn/examples/64_manager/gateway"
	"golang-learn/examples/64_manager/nsq"
	"golang-learn/examples/64_manager/process"
	"golang-learn/examples/64_manager/sqlcon"
	"golang-learn/examples/64_manager/web"
	"io"
	"os"
	"os/signal"
	"syscall"
)

// IManager ... manager interface
type IManager interface {
	// Processes
	GetProcess(key string) process.IProcess
	AddProcess(key string, process process.IProcess) error
	RemProcess(key string) (process.IProcess, error)

	// Configurations
	GetConfig(key string) config.IConfig
	AddConfig(key string, config config.IConfig) error
	RemConfig(key string) (config.IConfig, error)

	// Connections
	GetConnection(key string) (*sql.DB, error)
	AddConnection(key string, controller *sqlcon.SQLConController) error
	RemConnection(key string) (*sql.DB, error)

	// Gateways
	GetGateway(key string) (*gateway.Gateway, error)
	AddGateway(key string, gateway *gateway.Gateway) error
	RemGateway(key string) (*gateway.Gateway, error)
	RequestGateway(key string, method string, endpoint string, headers map[string]string, body io.Reader) (int, []byte, error)

	// NEW Instances
	NewSimpleConfig(path string, file string, extension string) (config.IConfig, error)
	NewSQLConnection(config *sqlcon.Config) (*sqlcon.SQLConController, error)
	NewNSQConsumer(config *nsq.Config, handler nsq.IHandler) (nsq.IConsumer, error)
	NewNSQProducer(config *nsq.Config) (nsq.IProducer, error)
	NewWEBServer(config *web.Config) (web.IWebController, error)
	NewGateway(config *gateway.Config) (*gateway.Gateway, error)

	// MANAGER
	Start() error
	Stop() error
}

// manager ... manager structure
type manager struct {
	ProcessController map[string]*process.ProcessController
	ConfigController  map[string]*config.ConfigController
	SqlConController  map[string]*sqlcon.SQLConController
	GatewayController map[string]*gateway.Gateway

	control chan int
	Started bool
}

// NewManager ... create a new manager
func NewManager() (IManager, error) {

	return &manager{
		ProcessController: make(map[string]*process.ProcessController),
		ConfigController:  make(map[string]*config.ConfigController),
		SqlConController:  make(map[string]*sqlcon.SQLConController),
		GatewayController: make(map[string]*gateway.Gateway),

		control: make(chan int),
	}, nil
}

// Start ... starts and blocks until it receives a signal in its control channel or a SIGTERM,
func (instance *manager) Start() error {
	log.Infof("manager, starting")
	instance.Started = true

	// listen for termination signals
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	// launch every process in a separeted process
	for name, process := range instance.ProcessController {
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
	if instance.Started {
		log.Infof("manager, stopping")

		for key, controller := range instance.ProcessController {
			if controller.Started {
				log.Infof("manager, stopping process [process:%s]", key)
				if err := controller.Process.Stop(); err != nil {
					log.Error(err, fmt.Sprintf("error stopping process [process:%s]", key))
				}
				log.Infof("manager, close channel [process:%s]", key)
				<-controller.Control
				close(controller.Control)
				delete(instance.ProcessController, key)
				log.Infof("manager, stopped process [process:%s]", key)
			}
		}

		instance.Started = false
		log.Infof("manager, stopped")
	}

	return nil
}
