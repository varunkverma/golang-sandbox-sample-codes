package main

import (
	"fmt"
	"errors"
)

type customerErr struct{
	name string
	age int
	err error 
}

func (c customerErr) Error() string{
	return fmt.Sprintf("%v has made some error\n",c.name)
}

func foo(e error){
	switch e.(type){
	case customerErr:
		fmt.Println("Customer error - ",e.(customerErr).err)
	default:
		fmt.Println(e.Error())
	}
}

func main(){
	cer:=customerErr{
		name:"John",
		age:34,
		err: errors.New("Customer has Erros"),
	}
	foo(cer)
}
