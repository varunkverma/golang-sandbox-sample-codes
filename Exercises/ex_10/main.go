package main

import "fmt"

const (
	y0=2016
	y1=y0+iota
	y2=y0+iota
	y3=y0+iota
	y4=y0+iota
)

func main(){
	fmt.Println(y1,y2,y3,y4)
}