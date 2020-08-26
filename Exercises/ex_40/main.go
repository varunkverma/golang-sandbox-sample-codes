package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){

	counter:=0

	goRoutines:=100
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i:=0;i<goRoutines;i++{
		go func(){
			v:=counter
			runtime.Gosched()
			v++
			counter=v
			wg.Done()
		}()
		fmt.Println("GoRoutine:",runtime.NumGoroutine())
		fmt.Println("Counter:",counter)
	}
	wg.Wait()
	fmt.Println("GoRoutine->:",runtime.NumGoroutine())
	fmt.Println("Counter->:",counter)
}