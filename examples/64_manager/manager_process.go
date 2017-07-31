package pm

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/nsq"
	"golang-learn/examples/64_manager/process"
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
func (instance *manager) GetProcess(key string) process.IProcess {
	return instance.ProcessController[key].Process
}

// AddProcess ... add a process with key
func (instance *manager) AddProcess(key string, prc process.IProcess) error {
	if instance.Started {
		panic("manager, can not add processes after start")
	}

	instance.ProcessController[key] = &process.ProcessController{
		Process: prc,
		Control: make(chan int),
	}
	log.Infof(fmt.Sprintf("manager, process '%s' added", key))

	return nil
}

// RemProcess ... remove the process by bey
func (instance *manager) RemProcess(key string) (process.IProcess, error) {
	// get process
	controller := instance.ProcessController[key]

	// delete process
	delete(instance.ProcessController, key)
	log.Infof(fmt.Sprintf("manager, process '%s' removed", key))

	return controller.Process, nil
}

// launch ... starts a process
func (instance *manager) launch(name string, controller *process.ProcessController) error {
	if err := controller.Process.Start(); err != nil {
		log.Error(err, fmt.Sprintf("manager, error launching process [process:%s]", name))
		instance.Stop()
		controller.Control <- 0
	} else {
		log.Infof(fmt.Sprintf("manager, launched process [process:%s]", name))
		controller.Started = true
		controller.Control <- 0
	}

	return nil
}
