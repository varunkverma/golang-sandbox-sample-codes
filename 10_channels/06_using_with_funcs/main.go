package main

import (
	"fmt"
)

// send
func foo(c chan<-int){
	c<-42
}
// receive
func bar(c <-chan int){
	fmt.Println(<-c)
}
func main(){
	// channel
	c:=make(chan int)

	// send
	go foo(c)

	//reveive
	bar(c)

	fmt.Println("Exiting")
}