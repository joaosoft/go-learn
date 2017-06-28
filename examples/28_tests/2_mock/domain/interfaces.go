package domain

type IRepository interface {
	DoSomething(id string, value string) error
}