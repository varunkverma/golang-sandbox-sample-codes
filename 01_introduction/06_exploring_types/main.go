package main

import "fmt"

var a="blah blah blah"

var b=`a said, "blah blah blah"`

var c=`a and
b
are idiots`

func main(){
	fmt.Printf("%T\n",a)
	fmt.Println(a)
	fmt.Printf("%T\n",b)
	fmt.Println(b)
	fmt.Printf("%T\n",c)
	fmt.Println(c)
}