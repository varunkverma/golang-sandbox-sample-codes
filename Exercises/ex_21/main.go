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

	people:=[]person{p1,p2}

	for _,per:=range people{
		fmt.Printf("%v %v likes:\n",per.firstName,per.lastName)
		for _,flavor:=range per.favIceCreams{
			fmt.Printf("%v\t",flavor)
		}
		fmt.Println()
	}
}