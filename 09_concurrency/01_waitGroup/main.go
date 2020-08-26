package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo(){
	for i:=0;i<10;i++{
		fmt.Printf("foo: %v\t",i)
	}
	fmt.Println()

	wg.Done()
}

func bar(){
	for i:=0;i<10;i++{
		fmt.Printf("bar: %v\t",i)
	}
	fmt.Println()
}

// func init(){
	// Runs only ones before main
// }

// func main is also a go-routine
func main(){

	fmt.Println("OS: ",runtime.GOOS)
	fmt.Println("ARCH: ",runtime.GOARCH)
	fmt.Println("CPUs: ",runtime.NumCPU())
	fmt.Println("Goroutines: ",runtime.NumGoroutine())
	
	wg.Add(1)
	// launch goroutine just by using keyword "go" / Concurrent design pattern
	go foo()
	bar()

	fmt.Println()
	fmt.Println("OS: ",runtime.GOOS)
	fmt.Println("ARCH: ",runtime.GOARCH)
	fmt.Println("CPUs: ",runtime.NumCPU())
	fmt.Println("Goroutines: ",runtime.NumGoroutine())
	wg.Wait()
}