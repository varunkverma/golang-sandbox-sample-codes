package main

import "fmt"

func main(){
	foo()
	numbers:= []int{1,2,3}
	fmt.Println("Before:",numbers)
	sliceChange(numbers)
	fmt.Println("After:",numbers)
	newNumbers :=sliceChange2(numbers) // returns [6/6]0xc000180060
	println("NEW NUMBERS: ",newNumbers)
	x,y:=greet()
	fmt.Println(x,y)
}

func foo(){
	fmt.Print("Foo here")
}

func sliceChange(numbers []int){
	fmt.Println("In func before:",numbers)
	numbers=append(numbers,4,5,6)
	fmt.Println("In func after:",numbers)
}

func sliceChange2(numbers []int) []int{
	fmt.Println("In func before:",numbers)
	numbers=append(numbers,4,5,6)
	fmt.Println("In func after:",numbers)
	return numbers
}

// return muliple values 
func greet() (string,int){
	return "Hello",101
}


// Note: When we pass values to a func. its called passing argsuments and when we receive values in a function, its called parameter received 

// Note: Everything in go is pass by value 