package process

// IProcess ... process interface
type IProcess interface {
	Start() error
	Stop() error
}

type IProcessController interface {
	Get() (IProcess, error)
	Start() error
	Stop() error
}

// processController ... controller structure
type ProcessController struct {
	Process IProcess
	Control chan int
	Started bool
}
