package main

import "fmt"

func main(){
	a:=incrementor()
	b:=incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int{
	x:=0
	return func() int{
		x++
		return x
	}
}