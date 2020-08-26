package main

import "fmt"

func foo(p *int){
	fmt.Println(*p)
	*p=45
	fmt.Println(*p)
}

func main(){
	x:=12
	fmt.Println("Before: ",x)
	foo(&x)
	fmt.Println("After: ",x)

}