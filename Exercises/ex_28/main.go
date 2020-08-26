package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func (p person) speak(){
	fmt.Println("first name:",p.first)
	fmt.Println("age:",p.age)
}

func main(){
	p:= person{
		first:"John",
		last:"Doe",
		age:16,
	}
	p.speak()
}