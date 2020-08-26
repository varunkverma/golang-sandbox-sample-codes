package main

import "fmt"

type person struct{
	firstName string
	lastName string
}

type secretAgent struct{
	person
	ltk bool
}

func (sa secretAgent) speak(){
	fmt.Printf("I'm %v, %v %v\n",sa.lastName,sa.firstName,sa.lastName)
}

// method speak gets attached to type person
func (p person) speak(){
	fmt.Printf("I'm the person -  %v %v\n",p.firstName,p.lastName)
}

// keyword indentifier type
// any type T which has speak() method attached to it will also be of type human
type human interface{
	speak()
}

// Note: every type implements an empty interface i.e., "interface{}", just like Objects in other languages

// polymorphism as this method can accept a type of type person or secretAgent
func bar(h human){
	// special type of switch statement T.(type), its called assertion
	switch h.(type){
	case person:
		fmt.Println("Person Type")
		fmt.Println(h.(person).firstName)
	case secretAgent:
		fmt.Println("Secret Agent Type")
		fmt.Println(h.(secretAgent).firstName)
	default:
		fmt.Printf("I'm of type %T\n",h)
	}
}

func main(){
	sa1:=secretAgent{
		person:person{
			firstName:"James",
			lastName:"Bond",
		},
		ltk:true,
	}
	p :=person{
		firstName:"John",
		lastName:"Doe",
	}

	bar(p)
	bar(sa1)
}
