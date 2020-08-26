package main

import "fmt"

func main(){
	a:=42
	fmt.Println(a) // value
	fmt.Println(&a) // referenece

	fmt.Printf("%T\n",a) // int
	fmt.Printf("%T\n",&a) // *int

	var ptrA *int=&a
	fmt.Println(ptrA)
	fmt.Println(*ptrA) // de-referencing an address
	fmt.Println(*&a)
	
	*ptrA=54
	fmt.Println(&a)
	fmt.Println(a)

}