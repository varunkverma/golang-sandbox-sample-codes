package main

import "fmt"

func main(){
	// declaring and initializing map
	m:= map[string]int {
		"Sal":2135469,
		"Harry":9978135,
	}
	fmt.Println(m)

	// access value through key
	fmt.Println(m["Sal"])

	// if key isn't present it returns the zero-value associated with the value
	a, ok:=m["Sal"]
	fmt.Println(a,ok)

	a,ok=m["Ron"]
	fmt.Println(a,ok)

	// use this in conditional control flow
	if a,ok=m["Harry"]; ok{
		fmt.Println("Calling Harry")
	}

	// use this in conditional control flow
	if a,ok=m["Sam"]; ok{
		fmt.Println("Calling Sam")
	}else{
		fmt.Println("Wrong number")
	}

	// declaring map using built-in method make
	n:=make(map[int]string)
	n[101]="Toys"
	n[102]="Books"
	fmt.Println(n)

	// iterate over map using range
	for k,v:=range n{
		fmt.Printf("%v -> %v\n",k,v)
	}

	// delete a key-value pair from map
	delete(n,101)
	fmt.Println(n)
	delete(n,101)

	// check before delete, although no error is through anyway
	if _,ok:=n[102];ok{
		delete(n,102)
	}
	fmt.Println(n)
}