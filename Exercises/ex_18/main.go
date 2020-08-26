package main

import "fmt"

func main(){
	arr:=[]int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2:7])
	fmt.Println(arr[1:6])

	arr=append(arr,52)
	fmt.Println(arr)
	arr=append(arr,53,54,55)
	fmt.Println(arr)
	x:=[]int{56,57,58,59,60}
	arr=append(arr,x...)
	fmt.Println(arr)
	y:=[2]int{61,62}
	//arr=append(arr,y...)// error
	fmt.Printf("%T\t%T\n",arr,y)
	arr2:=[]int{}
	arr2=append(arr2,arr[:3]...)
	arr2=append(arr2,arr[6:10]...)
	fmt.Println(arr2)
	
	arr3:=append(arr[:3],arr[6:10]...)
	fmt.Println(arr3)
}
