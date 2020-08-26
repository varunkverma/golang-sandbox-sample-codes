package main

import "fmt"

func mySum(xi... int) int {
	sum:=0
	for _,v:=range xi{
		sum+=v
	}
	return sum
}

func main(){
	fmt.Println("5+3 =",mySum(5,3))
}