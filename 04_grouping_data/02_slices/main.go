package main

import "fmt"

func main(){
	// composite literal
	x:=[]int{1,2,3,4,5,6,7}
	fmt.Println(x)

	// for range

	for i,v:=range x{
		fmt.Println(i,v)
	}

	fmt.Println(x[1:])
	//fmt.Println(x[1:100]) throws error
	fmt.Println(x[1:5])
	fmt.Println(x[1:1]) // empty slice

	x=append(x,8,9)
	fmt.Println(x)
	
	// spread operator
	y:=[]int{10,11,12,13,14}
	x=append(x,y...)
	fmt.Println(x)

	// allocation with build-in func main
	z:=make([]int,10,100) // make([]T,size,capacity)
	fmt.Println(z)
	fmt.Printf("Len: %v\n",len(z))
	fmt.Printf("Capacity: %v\n",cap(z))

	// multi-dimensional
	a:=[]string{"A","B","C"}
	b:=[]string{"E","F","G"}
	c:=[][]string{a,b}
	fmt.Println(c)
	fmt.Printf("%T",c)
}