package workers

type IController interface {
	Do() error
	Undo() error
}
