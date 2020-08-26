package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favIceCreams []string
}

func main(){
	p1:=person{
		firstName:"Jack",
		lastName:"Ryan",
		favIceCreams:[]string{"Chocolate","Vanilla"},
	}

	p2:=person{
		firstName:"Meg",
		lastName:"Ryan",
		favIceCreams:[]string{"Butterscotch","Vanilla"},
	}

	directory:=make(map[string]person)
	directory[p1.firstName]=p1
	directory[p2.firstName]=p2

	for k,v:=range directory{
		fmt.Printf("%v likes\n",k)
		for _,flavor:=range v.favIceCreams{
			fmt.Printf("%v ",flavor)
		}
		fmt.Println()
	}
}