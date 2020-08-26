package main

import "fmt"


var y=20 // package level scope
// z:=100, cannot use this short declaration operator outside the function body
var i int // assigns default value of type int to variable i
func main(){
	// Short declaration operator
	// declare a variable and assign a value (of a certain type)
	x:=42
	fmt.Println(x)
	fmt.Println(y)
	// fmt.Println(z)
	fmt.Println(i)
}