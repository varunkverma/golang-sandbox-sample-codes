package main

import "fmt"

func main(){
	// without switch parameter and fallthorugh by default is disabled
	switch{
	case false:
		fmt.Println("Case false");	
	case (2==5):
		fmt.Println("Case false");
	case (2==2):
		fmt.Println("Case True");
	case true:
		fmt.Println("Case True");
	}

	// fallthrough explicitly enabled
	switch{
	case false:
		fmt.Println("Case false");	
	case (2==5):
		fmt.Println("Case false");
	case (2==2):
		fmt.Println("Case True");
		fallthrough
	case true:
		fmt.Println("Case True");
		fallthrough
	case false==true:
		fmt.Println("False Case")
	}

	// fallthrough explicitly enabled
	switch{
	case false:
		fmt.Println("Case false");	
	case (2==5):
		fmt.Println("Case false");
	case (2==2):
		fmt.Println("Case True3");
		fallthrough
	case true:
		fmt.Println("Case True3");
		fallthrough
	default:
		fmt.Println("Default Case")
	}

	// multiple case options
	x:=5
	switch x{
	case 1,2,3:
		fmt.Println("1, 2 or 3")
	case 4,5,6:
		fmt.Println("4, 5 or 6")
	case 7,8,9:
		fmt.Println("7, 8 or 9")
	}

	// conditional case
	n:=5
	switch {
	case n%2==0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}
}