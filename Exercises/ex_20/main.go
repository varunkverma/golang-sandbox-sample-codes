package main

import "fmt"

func main(){
	favThings:=make(map[string][]string)
	favThings["James"]=[]string{"Shaken, not stired","Martinis","Women"}
	favThings["moneypenny"]=[]string{"James","Literature","CS"}
	favThings["no_dr"]=[]string{"Evil","Ice creeam","Sunsets"}
	fmt.Println(favThings)

	delete(favThings,"no_dr")

	for k,v :=range favThings{
		fmt.Printf("Fav things of %v\n",k)
		for _,item:=range v{
			fmt.Printf("%v\t",item)
		}
		fmt.Println()
	}
}