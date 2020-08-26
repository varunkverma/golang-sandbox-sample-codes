package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main(){

	goRoutines:=100
	
	var counter int64
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i:=0;i<goRoutines;i++{
		go func(){
			atomic.AddInt64(&counter,1)
			fmt.Println("GoRoutine:",runtime.NumGoroutine())
			fmt.Println("Counter:",atomic.LoadInt64(&counter))
			//runtime.Gosched()
			wg.Done()
		}()
		
	}
	wg.Wait()
	fmt.Println("GoRoutine->:",runtime.NumGoroutine())
	fmt.Println("Counter->:",counter)
}