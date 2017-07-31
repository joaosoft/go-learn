package pm

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/nsq"
)

// -------------- PROCESS CLIENTS --------------
// NewNSQConsumer ... creates a new nsq consumer
func (instance *manager) NewNSQConsumer(config *nsq.Config, handler nsq.IHandler) (nsq.IConsumer, error) {
	return nsq.NewConsumer(config, handler)
}

// NewNSQConsumer ... creates a new nsq producer
func (instance *manager) NewNSQProducer(config *nsq.Config) (nsq.IProducer, error) {
	return nsq.NewProducer(config)
}

// -------------- METHODS --------------
// GetProcess ... get a process with key
func (instance *manager) GetProcess(key string) IProcess {
	return instance.processController[key].process
}

// AddProcess ... add a process with key
func (instance *manager) AddProcess(key string, process IProcess) error {
	if instance.started {
		panic("manager, can not add processes after start")
	}

	instance.processController[key] = &ProcessController{
		process: process,
		control: make(chan int),
	}
	log.Infof(fmt.Sprintf("manager, process '%s' added", key))

	return nil
}

// RemProcess ... remove the process by bey
func (instance *manager) RemProcess(key string) (IProcess, error) {
	// get process
	controller := instance.processController[key]

	// delete process
	delete(instance.processController, key)
	log.Infof(fmt.Sprintf("manager, process '%s' removed", key))

	return controller.process, nil
}

// launch ... starts a process
func (instance *manager) launch(name string, controller *ProcessController) error {
	if err := controller.process.Start(); err != nil {
		log.Error(err, fmt.Sprintf("manager, error launching process [process:%s]", name))
		instance.Stop()
		controller.control <- 0
	} else {
		log.Infof(fmt.Sprintf("manager, launched process [process:%s]", name))
		controller.started = true
		controller.control <- 0
	}

	return nil
}
