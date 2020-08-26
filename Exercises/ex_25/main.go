package main

import "fmt"

func foo() int{
	return 1990
}

func bar() (int,string){
	return 18,"sd"
}

func main(){
	fmt.Println("foo:",foo())
	i,s:=bar()
	fmt.Println("bar: ",i,s)
}