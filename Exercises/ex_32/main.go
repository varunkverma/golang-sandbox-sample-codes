package main

import "fmt"

func sumOfNums(cb func(numbers ...int) int,nums ...int) int{
	return cb(nums...)
}

func main(){
	f:=func(numbers ...int) int{
		sum:=0
		for _,v:=range numbers{
			sum+=v
		}
		return sum
	}
	fmt.Println(sumOfNums(f,1,2,3,4))
}