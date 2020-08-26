package main

import (
	"fmt"
)

func send(e,o,q chan<-int){
	for i:=0;i<100;i++{
		if(i%2==0){
			e<-i
		}else{
			o<-i
		}
	}
	// close(e)
	// close(o)
	q<-0
	close(q)
}

func receive(e,o,q <-chan int){
	for{
		select{
		case v:=<-e:
			fmt.Println("even channel: ",v)
		case v:=<-o:
			fmt.Println("odd channel: ",v)
		case v,ok:=<-q:
			if(!ok){
				fmt.Println("From comma ok",v,ok)
				return
			}
			fmt.Println("From comma ok",v,ok)
			// close(q) // cannot close receive-only channel
			return
		}
	}
}

func main(){
	e:=make(chan int)
	o:=make(chan int)
	q:=make(chan int)

	// send
	go send(e,o,q)

	//receive
	receive(e,o,q)

	fmt.Println("Exiting")
}