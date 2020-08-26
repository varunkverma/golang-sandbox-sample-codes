package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
	password:="password123"
	bs,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	loginPassword1:="password123"
	errLogin:=bcrypt.CompareHashAndPassword(bs,[]byte(loginPassword1))
	if errLogin!=nil{
		fmt.Println(errLogin)
	}else{
		fmt.Println("Login successful")
	}

	loginPassword2:="password124"
	errLogin=bcrypt.CompareHashAndPassword(bs,[]byte(loginPassword2))
	if errLogin!=nil{
		fmt.Println(errLogin)
	}else{
		fmt.Println("Login successful")
	}
}