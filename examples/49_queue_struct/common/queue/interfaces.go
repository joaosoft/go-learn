package queue

type IController interface{
	Do() error
	Undo() error
}
