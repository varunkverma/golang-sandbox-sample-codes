package main

import "fmt"

func main(){
	
	fact:=func(n int)int{
		var f int=1
		for i:=2;i<=n;i++{
			f*=i
		}
		return f
	}(4)
	fmt.Println(fact)

	factExp:=func(n int)int{
		var f int=1
		for i:=2;i<=n;i++{
			f*=i
		}
		return f
	}
	fmt.Println("fact of 2",factExp(2))
	fmt.Println("fact of 3",factExp(3))
	fmt.Println("fact of 4",factExp(4))
	fmt.Println("fact of 5",factExp(5))
}