// THIS CODE DOESN"T WORK
package main

import (
	"fmt"
	//"sync"
)

func foo(c chan<-int){
	const goRoutines=10
	//var wg sync.WaitGroup
	for i:=0;i<goRoutines;i++{
		//wg.Add(i)
		go func(){
			for j:=0;j<10;j++{
				c<-j
			}
			//wg.Done()
		}()	
		//wg.Wait()
	}
	
	
}

func bar(c <-chan int){
	for{
		select{
		case v,ok:=<-c:
			if(ok){
				fmt.Printf("%v\t ok= %v\n",v,ok)
			}else{
				fmt.Printf("%v\t ok= %v\n",v,ok)
				return
			}
			
		}
	}
}

func main(){
	c:=make(chan int)

	foo(c)
	
	bar(c)

}