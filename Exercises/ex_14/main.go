package main

import (
	"fmt"
	"time"
)
func main(){
	birthYear:=1996
	currentYear:=time.Now().Year()
	for {
		if birthYear>currentYear{
			break;
		}
		fmt.Println(birthYear);
		birthYear++;
	}
}