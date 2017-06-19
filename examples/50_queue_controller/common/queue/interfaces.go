package queue

type IWork interface{
	GetWork() (interface{}, error)
}


type IController interface{
	Do(data interface{}) error
	Undo() error
}
