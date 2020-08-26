package main

import "fmt"

var x =42
var y ="James Bond"
var z =true

func main(){
	s:=fmt.Sprintln(x,y,z)
	fmt.Println(s)
	s2:=fmt.Sprintf("%v\t%v\t%v",x,y,z)
	fmt.Println(s2)
}