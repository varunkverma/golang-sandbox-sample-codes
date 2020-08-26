package main

import "fmt"

var y=42

func main(){
	fmt.Printf("%T\n",y)
	fmt.Printf("%b\n",y)
	fmt.Printf("%x\n",y)
	fmt.Printf("%#x\n",y)

	fmt.Printf("%#x\t%b\t%x\n",y,y,y)
	s:=fmt.Sprintf("%#x\t%b\t%x\n",y,y,y)
	fmt.Println(s)
	
}