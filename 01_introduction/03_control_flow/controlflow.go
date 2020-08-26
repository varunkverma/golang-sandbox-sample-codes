package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello this is control flow")
	foo()
	fmt.Println("Even nos")
	for i:=0;i<10;i++{
		if i%2 ==0{
			fmt.Println(i)
		}
	}
}

func foo(){
	n,err := fmt.Println("I'm foo")
	fmt.Println(n)
	fmt.Println(err)
}