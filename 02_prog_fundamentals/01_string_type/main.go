package main

import  "fmt"

func main(){
	s:="Hello World"
	fmt.Printf("%v\t%T\n",s,s)
	// converting string to collection of bytes
	bs:=[]byte(s)
	fmt.Printf("%v\t%T\n",bs,bs)

	for i:=0;i<len(s);i++{
		fmt.Printf("%#U",s[i])
	}

	fmt.Println()

	for i,v :=range s {
		fmt.Printf("index: %v\tvalue:%v\thex:%#x\n",i,v,v)
	}
}