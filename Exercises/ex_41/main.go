package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){

	counter:=0

	goRoutines:=100
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i:=0;i<goRoutines;i++{
		go func(){
			mutex.Lock()
			v:=counter
			v++
			counter=v
			fmt.Println("GoRoutine:",runtime.NumGoroutine())
			fmt.Println("Counter:",counter)
			mutex.Unlock()
			wg.Done()
		}()
		
	}
	wg.Wait()
	fmt.Println("GoRoutine->:",runtime.NumGoroutine())
	fmt.Println("Counter->:",counter)
}