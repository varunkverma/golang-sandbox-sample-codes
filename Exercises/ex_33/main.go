package main

import "fmt"

type person struct{
	first string
	last string 
	age int
}

func changeMe(p *person){
	(*p).first="John"
	(*p).last="Doe"
	p.age=42
}

func main(){
	a:=10
	fmt.Println(a,&a)

	p:= person{
		first:"Rose",
		last:"Bin",
		age:24,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}