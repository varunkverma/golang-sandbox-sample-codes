package main

import (
	"fmt"
	"log"
	"errors"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

// using a local errors.New()
func sqrt(f float64)(float64,error){
	if f<0{
		return 0,errors.New("norgate math: square root of negative number")
	}
	return 42,nil
}

// using a package level defined Error
func sqrt2(f float64)(float64,error){
	if f<0{
		return 0,ErrNorgateMath
	}
	return 42,nil
}

// using fmt.Errorf()
func sqrt3(f float64)(float64,error){
	if f<0{
		return 0,fmt.Errorf("This error is caused on receiving the parameter: %v",f)
	}
	return 42,nil
}

func sqrt4(f float64)(float64,error){
	if f<0{
		nme:=fmt.Errorf("This is a mathematical error :%v",f)
		return 0,norgateMathError{"34.134 N","99.0312 W",nme}
	}
	return 42,nil
}

// Note: Any type which implements a method of Error() string is also of type error
type norgateMathError struct{
	lat string
	long string
	err error
}

func (n norgateMathError) Error() string{
	return fmt.Sprintf("a Norgate error occured: %v %v %v",n.lat,n.long,n.err)
}


func main(){
	// _,err:=sqrt(-10)
	// if err != nil{
	// 	log.Fatalln(err)
	// }
	fmt.Printf("%T\n",ErrNorgateMath)
	_,err:=sqrt4(-10)
	if err != nil{
		log.Fatalln(err)
	}
}
