package main

import (
	"fmt"
)

func main(){

	// bi-directional channel
	c:=make(chan int ,2)
	
	c<-42
	c<-43 

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n",c)

	// directional-channel, send-only
	cw:=make(chan<-int,2)
	cw<-1
	cw<-2

	cw2:=make(chan<-int)
	go func(){
		cw2<-42
	}()
	// fmt.Println(<-cw2) // send-only channel
	
	// fmt.Println(<-cw) error 
	// fmt.Println(<-cw) error

	// directional-channel, receive-only - but error will still come
	 cr:=make(<-chan int,1)
	 fmt.Printf("%T\n",cr) 
	// fmt.Println(<-cr) //fatal error: all goroutines are asleep - deadlock!

	// specific to general doesn't assign
	// c=cw //error
	// c=cr //error

	// general to specific assigns
	cw=c
	cr=c
	fmt.Printf("%T\n",cr)

	// general to specific converts
	fmt.Printf("%T\n",(<-chan int)(c))
	fmt.Printf("%T\n",(chan<-int)(c))

	// specific to general doesn't convert
	
}