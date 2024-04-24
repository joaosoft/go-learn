package worker

import (
	"github.com/joaosoft/golang-learn/46_workers/controllers"
)

type WorkerController struct {
	SomeController controllers.SomeController
}

// NewWorkerController ...
func NewWorkerController(controller controllers.SomeController) *WorkerController {
	workerController := new(WorkerController)
	workerController.SomeController = controller

	return workerController
}

// Do ...
func (workerController *WorkerController) Do() error {
	return workerController.SomeController.DoSomething()
}

// Undo ...
func (workerController *WorkerController) Undo() error {
	// Do Something ...
	return nil
}
