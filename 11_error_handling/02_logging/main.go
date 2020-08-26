package main

import (
	"fmt"
	"log"
	"os"
)

func simpleLogging(){
	f,err:=os.Create("log.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer f.Close()
	// Setting the log storage destination
	log.SetOutput(f)

	f2,err:=os.Open("no-file.txt")
	if err!=nil{
		log.Println("error happened",err)
	}
	defer f2.Close()
	fmt.Println("Check the log.txt file in the directory")
}

func fatalFunc(){
	f,err:=os.Open("no-file.txt")
	f.Close()
	if err!=nil{
		// the Fatal functions call os.Exit(1) after writing the log message. They don't wait for deferred functions
		log.Fatalln(err)
	}
}

func main(){
	fmt.Println("Before Fatal")
	fatalFunc()
	fmt.Println("After Fatal") // will not be reached and printed due to Fatalln func
}