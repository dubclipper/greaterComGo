package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) details() {
	fmt.Printf("%s\n", e.name)
	fmt.Printf("%d\n", e.age)
}

func (e employee) getSalary() int {
	return e.salary
}

func main() {
	emp := employee{
		name:   "Jane",
		age:    33,
		salary: 38392,
	}
	emp.details()
	fmt.Printf("Salary %d\n", emp.getSalary())
}
