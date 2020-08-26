package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	//receive1(c, q)
	receive2(c, q)

	fmt.Println("about to exit")
}
func receive1(c <-chan int, q chan int){
	for{
		select{
		case v:=<-c:
			fmt.Println("from channel c:",v)
		case v,ok:=<-q:
			if(!ok){
				fmt.Println("from comma ok q:",v,ok)
				return
			}
			fmt.Println("from comma ok q:",v,ok)
			return
		}
	}
}
func receive2(c <-chan int, q chan int){
	for{
		select{
		case v:=<-c:
			fmt.Println("from channel c:",v)
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func(){
		for i := 0; i < 100; i++ {
			c <- i
		}
		q<-1
		close(c)
	}()

	return c
}
