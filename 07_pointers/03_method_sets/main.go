package main

import "fmt"

// non-pointer receiver: Can receive non-pointer Type values
func foo(a int){
	fmt.Println(a)
}
func bar(a *int){
	fmt.Println(a)
	fmt.Println(*a)
}

func main(){
	a:=10
	foo(a) // passing a non-pointer type value
	//foo(&a) // passing a pointer type; error
	bar(&a)
	//bar(a) // error
}