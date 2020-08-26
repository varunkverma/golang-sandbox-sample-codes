package main

import (
	"fmt"
	"encoding/json"
)

type user struct{
	Username string `json:"Username"`
	Role string	`json:"Role"`
}

func main(){
	u1:=user{
		Username:"JohnDoe",
		Role:"Admin",
	}
	u2:=user{
		Username:"JaneDoe",
		Role:"PremiumUser",
	}

	users:=[]user{
		u1,
		u2,
	}

	bs,err:=json.Marshal(users)
	if err!=nil{
		fmt.Println(err)
	}
	
	jsonString:=string(bs)
	
	fmt.Println("users: ",jsonString)

	unmarshaledUsers:=[]user{}

	err2:=json.Unmarshal([]byte(jsonString),&unmarshaledUsers)

	if err2!=nil{
		fmt.Println(err2)
	}
	fmt.Println(unmarshaledUsers)
	for _,v:=range unmarshaledUsers{
		fmt.Println(v)
	}

}