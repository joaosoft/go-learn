package repositories

type IRepository interface {
	DoSomething(data interface{}) error
}
