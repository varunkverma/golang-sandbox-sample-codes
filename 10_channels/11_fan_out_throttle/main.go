package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func populate(c chan int){
	for i:=0;i<100;i++{
		c<-i
	}
	close(c)
}

func fanOutIn(c1,c2 chan int){
	var wg sync.WaitGroup
	// For throttling, we just are just gonna use a pre-defined limitted go-routines
	const goRountines=10
	wg.Add(goRountines)
	for i:=0;i<goRountines;i++{
		go func(){
			for v:=range c1{
				func(v2 int){
					c2<-timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n+rand.Intn(100)
}

func main(){
	c1:=make(chan int)
	c2:=make(chan int)

	go populate(c1)

	go fanOutIn(c1,c2)

	for v:=range c2{
		fmt.Println(v)
	}

	fmt.Println("Exiting the main")
}

