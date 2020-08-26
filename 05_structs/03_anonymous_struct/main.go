package main

import "fmt"

func main(){
	// anonymous struct
	p:=struct{
		first string
		last string
	}{
		first:"James",
		last:"Bond",
	}

	fmt.Println(p)
	fmt.Printf("%T\n",p)
}