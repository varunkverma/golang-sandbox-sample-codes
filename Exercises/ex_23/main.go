package main

import "fmt"

type vehicle struct{
	doors int
	color string
}

type truck struct{
	vehicle
	fourWheel bool
}

type sedan struct{
	vehicle
	luxury bool
}

func main(){
	tata:= truck{
		vehicle:vehicle{
			doors:4,
			color:"Red",
		},
		fourWheel:true,
	}
	honda:= sedan{
		vehicle:vehicle{
			doors:4,
			color:"Blue",
		},
		luxury:true,
	}
	fmt.Println(tata)
	fmt.Println(honda)
	fmt.Println(tata.vehicle.color)
}