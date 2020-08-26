package main

import (
	"fmt"
	"sort"
)

type byInt []int

func (a byInt) Len() int{
	return len(a)
}

func (a byInt) Less(i,j int) bool{
	return a[i]<a[j]
}

func (a byInt) Swap(i,j int){
	a[i],a[j]=a[j],a[i]
}

type byString []string

func (a byString) Len() int{
	return len(a)
}

func (a byString) Less(i,j int) bool{
	return a[i]<a[j]
}

func (a byString) Swap(i,j int){
	a[i],a[j]=a[j],a[i]
}


func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	//sort.Sort(byInt(xi))
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	//sort.Sort(byString(xs))
	fmt.Println(xs)
}
