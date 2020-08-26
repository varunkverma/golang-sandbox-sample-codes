package main

import "fmt"

type exty int
var x exty

func main(){
	fmt.Println(x)
	fmt.Printf("%v has a type of: %T\n",x,x)
	x=42
	fmt.Println(x)
}