// error is an interface, any type which implements Error(), is a string
// always check your errors

package main

import ( 
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"strings"
)

func main1(){
	var answer1,answer2,answer3 string

	fmt.Print("Name: ")
	_,err:=fmt.Scan(&answer1)
	if err!=nil{
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_,err=fmt.Scan(&answer2)
	if err!=nil{
		panic(err)
	}

	fmt.Print("Fav Color: ")
	_,err=fmt.Scan(&answer3)
	if err!=nil{
		panic(err)
	}

	fmt.Println(answer1,answer2,answer3)
}

// Writing to a file
func main2(){
	// create a new txt file
	f,err:=os.Create("names.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}
	// close the file before exiting the code, but defer it for now
	defer f.Close()

	// Reader
	r:=strings.NewReader("John Doe")

	// copy the string from Reader into the file
	io.Copy(f,r)

}

//Reading from a file
func main(){
	f,err:=os.Open("names.txt")
	if(err!=nil){
		fmt.Println(err)
		// panic(err), throws an error and stops further execution of the program
		return
	}
	defer f.Close()

	bs,err:=ioutil.ReadAll(f)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}