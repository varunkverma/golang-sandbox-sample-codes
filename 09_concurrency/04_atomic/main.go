package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main(){
	fmt.Println("CPUs: ",runtime.NumCPU())
	fmt.Println("Goroutines: ",runtime.NumGoroutine())

	var counter int64 // int64, a clue for using atomic package to prevent race condition

	const goRoutines=100
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i:=0;i<goRoutines;i++{
		go func(){
			atomic.AddInt64(&counter,1) // write to counter safely
			fmt.Println("Counter: ",atomic.LoadInt64(&counter)) // read from counter safely
			// time.Sleep(time.Second)
			runtime.Gosched() //Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			wg.Done()
		}()
		fmt.Println("Goroutines: ",runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines: ",runtime.NumGoroutine())
	fmt.Println("counter: ",counter)

}

// Note: Check if the code contains a race condition using the command " go run -race main.go "