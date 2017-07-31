package pm

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
	process IProcess
	control chan int
	started bool
}
