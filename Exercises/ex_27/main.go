package main

import "fmt"

func foo(){
	fmt.Println("Entering foo")
	bar()
	fmt.Println("Exiting foo")
}

func bar() {
	fmt.Println("Bar")
}

func main(){
	foo()
}


// Note: Cannot use defer wiht a function returnig something