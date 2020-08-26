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

	// mutex prevents race condition
	var mutex sync.Mutex

	for i:=0;i<goRoutines;i++{
		go func(){
			// locking the access to the variable to prevent other goroutines from accessing them
			mutex.Lock()
			v:=counter
			// time.Sleep(time.Second)
			// runtime.Gosched() //Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically. But no need if mutex is used
			v++
			counter=v
			// unlocking the access to the varibale so that other goroutine can access it
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines: ",runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines: ",runtime.NumGoroutine())
	fmt.Println("counter: ",counter)

}