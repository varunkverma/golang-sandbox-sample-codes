package main

import "fmt"
import "math"

type square struct{
	side float64
}

type circle struct{
	radius float64 
}

func (s square) area() float64{
	return s.side*s.side
}

func (c circle) area() float64{
	return math.Pi*(c.radius * c.radius)
}

type shape interface{
	area() float64
}

func info(s shape){
	var res float64
	switch s.(type){
	case square:
		res=s.(square).area()
	case circle:
		res=s.(circle).area()
	default:
		res=0
	}
	
	fmt.Println(res)
}

func main(){
	sq:=square{
		side:10,
	}

	ci:=circle{
		radius:1.7,
	}

	fmt.Println("Area of square:",sq.area())
	fmt.Println("Area of circle:",ci.area())

	info(sq)
	info(ci)
}