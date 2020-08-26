package main

import "fmt"

func main(){
	student:=struct{
		marks map[string]int
		friends []string
	}{
		marks:map[string]int{
			"Physics":80,
			"Chemistry":78,
			"Computer Science":95,
		},
		friends:[]string{"Meg,Rachel,Joey"},
	}
	fmt.Println(student)
	for subject,score:=range student.marks{
		fmt.Printf("%v -> %v\n",subject,score)
	} 
	for _,friend:=range student.friends{
		fmt.Println(friend)
	}
	
}