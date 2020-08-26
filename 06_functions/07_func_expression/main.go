package main

import "fmt"

func bar() func() int{
	return func() int{
		return 199
	}
}

func main(){
	// functions are first-class citizen
	f:=func(){
		fmt.Println("func expression")
	}
	f()
	g:=func(x int){
		fmt.Println(x)
	}
	g(545)

	numFunc:=bar()
	r:=numFunc()
	fmt.Println(r)
	fmt.Printf("%T\n",bar)
	fmt.Printf("%T\n",numFunc)

	fmt.Println(bar()())
}