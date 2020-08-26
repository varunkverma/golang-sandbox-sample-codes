package main

import (
	"fmt"
	"time"
)
func main(){
	birthYear:=1996
	currentYear:=time.Now().Year()
	for i:=birthYear;i<=currentYear;i++{
		fmt.Println(i)
	}
}