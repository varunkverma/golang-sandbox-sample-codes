package main

import (
	"fmt"
)

func foo(c chan<-int){
	go func(){
		for i:=0;i<100;i++{
			c<-i
		}
		close(c)
	}()
}

func bar(c <-chan int){
	for{
		select{
		case v,ok:=<-c:
			if(ok){
				fmt.Printf("%v\t ok= %v\n",v,ok)
			}else{
				fmt.Printf("%v\t ok= %v\n",v,ok)
				return
			}
			
		}
	}
}

func main(){
	c:=make(chan int)

	foo(c)

	bar(c)

}