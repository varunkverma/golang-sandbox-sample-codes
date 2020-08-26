package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T){
	c:=Count("Hello World")
	if c!=2{
		t.Error("Expected:",2,"Got:",c)
	}
}

func ExampleCount(){
	fmt.Println(Count("Hello World"))
	// Output:
	// 2
}

func BenchmarkCount(b *testing.B){
	for i:=0;i<b.N;i++{
		Count("Hello World")
	}
}

func BenchmarkUseCount(b *testing.B){
	for i:=0;i<b.N;i++{
		UseCount("Hello World")
	}
}