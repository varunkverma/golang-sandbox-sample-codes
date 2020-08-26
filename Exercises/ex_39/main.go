package main

import (
	"fmt"
)

type person struct{
	name string
	age int
}

func (p *person) speak(){
	fmt.Println("Hello my name is",p.name,", I'm",p.age,"years old")
}

type human interface{
	speak()
}

func saySomething(h human){
	h.speak()
}

func main(){
	p:=person{
		name:"John Doe",
		age:26,
	}
	ptrP:=&p

	saySomething(ptrP)
	// saySomething(p) // error
}