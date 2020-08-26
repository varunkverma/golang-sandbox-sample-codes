package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo(){
	fmt.Println("foo here")
	wg.Done()
}

func bar(){
	fmt.Println("bar here")
	wg.Done()
}

func main(){
	fmt.Println("Enter main")
	fmt.Println("GoRoutines:",runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("GoRoutines:",runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Exit main")
}