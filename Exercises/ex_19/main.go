package main

import "fmt"

func main(){
	names:=[][]string{
		{"James","Bond","Shaken, not stired"},
		{"Miss","Moneypenny","Hello James"},
	}
	for i,row:=range names{
		fmt.Println(i,row)
		for j,col:=range row{
			fmt.Printf("%v\t%v\t\t",j,col)
		}
		fmt.Println()
	}
}