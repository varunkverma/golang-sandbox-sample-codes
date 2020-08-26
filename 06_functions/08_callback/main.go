package main

import "fmt"

func main(){
	allNumbers:=[]int{1,2,3,4,5,6,7,8,9}
	sumOfEven:=sumOfEvenNums(sumOfNums,allNumbers...)
	fmt.Println("Sum of Even Numbers: ",sumOfEven) 
}

func sumOfNums(numbers ...int) int{
	sum:=0
	for _,v:=range numbers{
		sum+=v
	}
	return sum
}

func sumOfEvenNums(cb func(numbers ...int) int, allNumbers ...int) int{
	evenNums:=[]int{}
	for _,v:=range allNumbers{
		if v%2 ==0{
			evenNums=append(evenNums,v)
		}
	}
	return cb(evenNums...)
}