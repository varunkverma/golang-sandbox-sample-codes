package main

import "fmt"

func main(){
	// for loop
	for i:=0;i<=10;i++ {
		fmt.Println(i)
	}

	// nested loop
	for i:=0;i<5;i++{
		for j:=5-i+1;j<=5;j++{
			fmt.Print("*")
		}
		fmt.Println();
	}
	for i:=0;i<5;i++{
		for j:=i+1;j<=5;j++{
			fmt.Print("*")
		}
		fmt.Println();
	}

	// while loop
	x:=1
	for x<=10{
		fmt.Println(x)
		x++
	}

	// for-break
	y:=1
	for{
		if y==5{
			break;
		}
		fmt.Println(y)
		y++
	} 

	// continue keyword
	for i:=0;i<=10;i++{
		if i%2!=0{
			continue;
		}
		fmt.Println(i)
	}
}