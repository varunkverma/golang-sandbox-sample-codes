package main

import "fmt"

func main(){
	func(){
		fmt.Println("anonymous func")
	}()

	func(x int){
		fmt.Println(x)
	}(10)
}