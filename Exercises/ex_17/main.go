package main

import "fmt"

func main(){
	arr:=[]int{1,2,3,4,5}
	for i,v:= range arr{
		fmt.Printf("index: %v -> value: %v\n",i,v)
	}
	fmt.Printf("%T\n",arr)
}