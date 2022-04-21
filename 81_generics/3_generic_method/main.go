package main

import "fmt"

type Employee interface {
	PrintSalary()
}

func getSalary[E Employee](e E) {
	e.PrintSalary()
}

type Engineer struct {
	Salary int32
}

func (e Engineer) PrintSalary() {
	fmt.Println(e.Salary)
}

type Manager struct {
	Salary int64
}

func (m Manager) PrintSalary() {
	fmt.Println(m.Salary)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	engineer := Engineer{Salary: 10}
	manager := Manager{Salary: 100}

	getSalary(engineer)
	getSalary(manager)
}
