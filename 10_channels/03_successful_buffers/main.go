package main

import (
	"fmt"
)

func main(){
	
	// buffered channel, size is mentioned while declaring a channel
	c:=make(chan int ,1)
	
	c <- 42

	fmt.Println(<-c)
}