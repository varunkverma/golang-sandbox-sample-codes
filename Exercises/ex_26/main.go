package main

import "fmt"

func foo(nums ...int) int{
	sum:=0
	for _,v:=range nums{
		sum+=v
	}
	return sum
}

func bar(nums []int) int{
	sum:=0
	for _,v:=range nums{
		sum+=v
	}
	return sum
}

func main(){
	fmt.Println("foo:",foo(1,2,3,4,5,6))
	fmt.Println("bar:",bar([]int{1,2,3,4,5,6}))
}