package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	// using func literal aka anonymous self-executong func
	go func(){
		c<-42
	}()
	
	// using buffered channel
	c2:=make(chan int,1)
	c2<-42

	fmt.Println(<-c)
	fmt.Println(<-c2)
}
