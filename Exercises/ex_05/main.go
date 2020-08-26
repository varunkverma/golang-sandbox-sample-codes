package main

import "fmt"

type exty int
var x exty

type exty2 exty
var y exty2

func main(){
	fmt.Println(x)
	fmt.Printf("%v has a type of: %T\n",x,x)
	x=42
	fmt.Printf("%T\n",x)
	//y=x // error, though y:=x, will re-declares the variable y
	y:=exty2(x)
	fmt.Printf("%v\t%T",y,y)
}