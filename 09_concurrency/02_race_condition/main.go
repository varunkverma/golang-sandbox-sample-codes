package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	fmt.Println("CPUs: ",runtime.NumCPU())
	fmt.Println("Goroutines: ",runtime.NumGoroutine())

	counter:=0

	const goRoutines=100
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i:=0;i<goRoutines;i++{
		go func(){
			v:=counter
			// time.Sleep(time.Second)
			runtime.Gosched() //Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			v++
			counter=v
			wg.Done()
		}()
		fmt.Println("Goroutines: ",runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines: ",runtime.NumGoroutine())
	fmt.Println("counter: ",counter)

}

// Note: Check if the code contains a race condition using the command " go run -race main.go "