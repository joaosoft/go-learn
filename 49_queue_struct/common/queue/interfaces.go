package queue

type IController interface {
	Do(data []byte) error
}

type IWork interface {
	GetWork() ([]byte, error)
}
