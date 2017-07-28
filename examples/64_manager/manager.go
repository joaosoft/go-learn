package pm

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/nsq"
	"os"
	"os/signal"
	"syscall"
)

// IManager ... manager interface
type IManager interface {
	GetProcess(key string) IProcess
	AddProcess(key string, process IProcess) error
	RemProcess(key string) (IProcess, error)

	GetConfig(key string) IConfig
	AddConfig(key string, config IConfig) error
	RemConfig(key string) (IConfig, error)

	NewNSQConsumer(config *nsq.Config, handler nsq.IHandler) (nsq.IConsumer, error)
	NewNSQProducer(config *nsq.Config) (nsq.IProducer, error)

	NewSimpleConfig(path string, file string, extension string) (IConfig, error)

	Start() error
	Stop() error
}

// manager ... manager structure
type manager struct {
	processController map[string]*processController
	configController  map[string]*configController

	control chan int
	started bool
}

// NewManager ... create a new manager
func NewManager() IManager {
	return &manager{
		processController: make(map[string]*processController),
		configController:  make(map[string]*configController),
		control:           make(chan int),
	}
}

// Start ... starts and blocks until it receives a signal in its control channel or a SIGTERM,
func (instance *manager) Start() error {
	log.Debug("manager, starting")
	instance.started = true

	// listen for termination signals
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	// launch every process in a separeted process
	for name, process := range instance.processController {
		log.Infof("manager, starting process [process:%s]", name)

		go instance.launch(name, process)

		log.Debugf("manager, started process [process:%s]", name)
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
		log.Debugf("manager, stopping")

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
