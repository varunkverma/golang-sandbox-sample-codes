package main

import "fmt"

type person struct{
	first string
	last string
}

type secretAgent struct{
	person // anonymous field
	licsenseToKill bool // named field
}

func main(){
	p1:=person{
		first:"James",
		last:"Bond",
	}
	p2:=person{
		first:"Miss",
		last:"MoneyPenny",
	}

	sa1:=secretAgent{
		person:p1,
		licsenseToKill:true,
	}
	sa2:=secretAgent{
		person:p2,
		licsenseToKill:false,
	}

	fmt.Println(p1.last)
	fmt.Println(p2.first)

	fmt.Println(sa1)
	fmt.Println(sa2)

	fmt.Println(sa1.person.first)
	// or (unless you have a same key in person and secretAgent)
	fmt.Println(sa1.first)
}