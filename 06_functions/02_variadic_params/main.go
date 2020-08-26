package main

import "fmt"

func main(){
	foo(1,2,3,4,5,6,7,8,9,10)
	s:=add(1,2,3,4,5)
	fmt.Println(s)

	nums :=[]int{1,2,3,4,5,6}
	s2:=add(nums...)
	fmt.Println(s2)

	s3:=add()
	fmt.Println(s3)
}

func foo(x...int){ //unfirling
	fmt.Println(x)
	fmt.Printf("%T\n",x) // slice of int
}

func add(numbers ...int)int{
	sum:=0
	for _,v:=range numbers{
		sum+=v
	}
	return sum
}

