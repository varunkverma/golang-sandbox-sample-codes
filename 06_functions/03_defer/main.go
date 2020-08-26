package main

import "fmt"

func main(){
	defer foo()
	bar()
}

func foo(){
	fmt.Println("foo")
}

func bar(){
	defer fmt.Println("bar")
	fmt.Println("wait for me")
}