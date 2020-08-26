package main

import (
	"fmt"
)

func main(){

	// buffered channel, size is mentioned while declaring a channel
	c:=make(chan int ,1)
	
	c<-42
	c<-43 // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-c)
}