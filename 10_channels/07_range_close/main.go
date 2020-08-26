package main

import (
	"fmt"
)

func main(){
	// channel
	c:=make(chan int)

	// send
	go func(){
		for i:=0;i<100;i++{
			c<-i
		}
		close(c)
	}()

	//reveive
	for v:=range c{
		fmt.Printf("%v\t",v)
	}

	fmt.Println("Exiting")
}