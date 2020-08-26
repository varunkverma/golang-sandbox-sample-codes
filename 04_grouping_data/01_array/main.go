package main

import "fmt"

func main(){
	var x [5]int;
	fmt.Println(x)
	x[2]=12
	fmt.Println(x,len(x))
	// better to use slices

	xa:=[]int{1,2,3,4,5,6,7}
	fmt.Println(xa)
}