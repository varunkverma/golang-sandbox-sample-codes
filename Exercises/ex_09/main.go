package main

import "fmt"

func main(){
	x:=10
	fmt.Printf("%d\t%b\t%#x\n",x,x,x)
	y:=x<<1
	fmt.Printf("%d\t%b\t%#x\n",y,y,y)
	z:=x>>1
	fmt.Printf("%d\t%b\t%#x\n",z,z,z)
	z=z>>1
	fmt.Printf("%d\t%b\t%#x\n",z,z,z)
}