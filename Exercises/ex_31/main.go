package main

import "fmt"

func foo() func(n int) int{
	return func(n int) int{
		var f int=1
		for i:=2;i<=n;i++{
			f*=i
		}
		return f
	}
}

func main(){
	factFunc:=foo()
	fmt.Println(factFunc(3))
}