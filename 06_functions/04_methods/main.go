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

// Method declaration
// syntax:
// func (r receiver) identifier(parameters) (return(s)) {code...}
// it attaches the function to the receiver type
func (sa secretAgent) speak(){
	fmt.Printf("I'm %v, %v %v\n",sa.lastName,sa.firstName,sa.lastName)
}

func main(){
	sa1:=secretAgent{
		person:person{
			firstName:"James",
			lastName:"Bond",
		},
		ltk:true,
	}

	fmt.Println(sa1)
	sa1.speak()
}