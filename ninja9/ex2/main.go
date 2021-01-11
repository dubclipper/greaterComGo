package main

import (
	"fmt"
)

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak(){
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "James"
	}
	saySomething(&p1)
	p1.speak()
}