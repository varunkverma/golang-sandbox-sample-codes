package main

import "fmt"

var a int

type hotdog int
var b hotdog

func main(){
	a=12
	b=98
	fmt.Println(a)
	fmt.Printf("%T\n",b)
	fmt.Println(b)
	fmt.Println()
	fmt.Println("CONVERSION:")
	a=int(b) // Conversion
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	

}