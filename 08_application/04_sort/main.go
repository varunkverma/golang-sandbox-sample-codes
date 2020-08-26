package main

import (
	"fmt"
	"sort"
)

type person struct{
	first string
	age int
}

type byAge []person

// functions necessary for sort interface
func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i,j int){
	temp:=a[i]
	a[i]=a[j]
	a[j]=temp
}

func (a byAge) Less(i,j int) bool{
	return a[i].age<a[j].age
}

type byName[]person

// functions necessary for sort interface
func (a byName) Len() int {
	return len(a)
}

func (a byName) Swap(i,j int){
	// shorter way of swap
	a[i],a[j]=a[j],a[i]
}

func (a byName) Less(i,j int) bool{
	return a[i].first<a[j].first
}

func main(){
	nums:=[]int{6,2,5,1,3,1,88,2,5}
	names:=[]string{"A","b","a","sadasr","qwe"}
	sort.Ints(nums)
	fmt.Println(nums)
	sort.Strings(names)
	fmt.Println(names)

	p1:=person{"James",32}
	p2:=person{"Harry",18}
	p3:=person{"Emma",25}
	p4:=person{"Mishael",35}

	people:=[]person{
		p1,p2,p3,p4,
	}
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
	fmt.Println()
	sort.Sort(byName(people))
	fmt.Println(people)
}