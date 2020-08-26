package main

import(
	"fmt"
	"encoding/json"
)

type person struct{
	First string `json:"First"`
	Last string `json:"Last"`
	Age int `json:"Age"`
}

func main(){
	s:=`[{"First":"Dave","Last":"Rasen","Age":30},{"First":"Kail","Last":"Summer","Age":36}]`
	bs:=[]byte(s)
	fmt.Printf("%T\n",s)
	fmt.Printf("%T\n",bs)

	people:= []person{}

	err:=json.Unmarshal(bs,&people)

	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("data: ",people)
	for _,v:=range people{
		fmt.Println(v.First,v.Last)
	}
}