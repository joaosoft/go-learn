package domain

type IRepository interface {
	Store(name string, age int) error
}
