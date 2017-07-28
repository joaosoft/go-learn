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
	GetProcess(key string) (IProcess, error)
	AddProcess(key string, process IProcess) error
	RemProcess(key string) (IProcess, error)
	Start() error
	Stop() error
	NewNSQConsumer(config *nsq.Config, handler nsq.IHandler) (nsq.IConsumer, error)
	NewNSQProducer(config *nsq.Config) (nsq.IProducer, error)
}

// manager ... manager structure
type manager struct {
	processProcessController map[string]ProcessController
	control                  chan int
	started                  bool
}

// NewManager ... create a new manager
func NewManager() IManager {
	return &manager{
		processProcessController: make(map[string]ProcessController),
		control:                  make(chan int),
	}
}

// ProcessController ... controller structure
type ProcessController struct {
	process IProcess
	control chan int
	started bool
}

// GetProcess ... get a process with key
func (instance *manager) GetProcess(key string) (IProcess, error) {
	return instance.processProcessController[key].process, nil
}

// AddProcess ... add a process with key
func (instance *manager) AddProcess(key string, process IProcess) error {
	if instance.started {
		panic("manager, can not add processes after start")
	}

	instance.processProcessController[key] = ProcessController{
		process: process,
		control: make(chan int),
	}
	log.Debug(fmt.Sprintf("manager, process '%s' added", key))

	return nil
}

// RemProcess ... remove the process by bey
func (instance *manager) RemProcess(key string) (IProcess, error) {
	// get process
	controller := instance.processProcessController[key]

	// delete process
	delete(instance.processProcessController, key)
	log.Debug(fmt.Sprintf("manager, process '%s' removed", key))

	return controller.process, nil
}

// Start ... starts and blocks until it receives a signal in its control channel or a SIGTERM,
func (instance *manager) Start() error {
	log.Debug("manager, starting")
	instance.started = true

	// listen for termination signals
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	// launch every process in a separeted process
	for name, process := range instance.processProcessController {
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

		for key, controller := range instance.processProcessController {
			log.Infof("manager, stopping process [process:%s]", key)
			if err := controller.process.Stop(); err != nil {
				log.Error(err, fmt.Sprintf("error stopping process [process:%s]", key))
			}
			log.Infof("manager, waiting for process to terminate [process:%s]", key)
			<-controller.control
			close(controller.control)
			delete(instance.processProcessController, key)
			log.Infof("manager, stopped process [process:%s]", key)
		}

		instance.started = false
		log.Infof("manager, stopped")
	}

	return nil
}

// NewNSQConsumer ... creates a new nsq consumer
func (instance *manager) NewNSQConsumer(config *nsq.Config, handler nsq.IHandler) (nsq.IConsumer, error) {
	return nsq.NewConsumer(config, handler)
}

// NewNSQConsumer ... creates a new nsq producer
func (instance *manager) NewNSQProducer(config *nsq.Config) (nsq.IProducer, error) {
	return nsq.NewProducer(config)
}

// launch ... starts a process
func (instance *manager) launch(name string, controller ProcessController) error {
	if err := controller.process.Start(); err != nil {
		log.Error(err, fmt.Sprintf("manager, error launching process [process:%s]", name))
		instance.Stop()
		controller.control <- 0
	} else {
		log.Info(fmt.Sprintf("manager, launched process [process:%s]", name))
		controller.started = true
		controller.control <- 0
	}

	return nil
}
