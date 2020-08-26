package main

import (
	"fmt"
	"encoding/json"
)

type person struct{
	// note: fields starting with lowercase represents private field and if starts with uppercase its a public field 
	First string
	Last string
	Age int
}

func main(){
	p1:=person{
		First:"Dave",
		Last:"Rasen",
		Age:30,
	}
	p2:=person{
		First:"Kail",
		Last:"Summer",
		Age:36,
	}
	people:=[]person{
		p1,
		p2,
	}

	fmt.Println(people)

	bs,err:=json.Marshal(people)
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
}