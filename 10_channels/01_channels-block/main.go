package main

import "fmt"

func main(){
	// create a channel to store int value, but this program results in deadlock
	//fatal error: all goroutines are asleep - deadlock!

	c:=make(chan int)
	c <- 42
	fmt.Println(<-c)
}