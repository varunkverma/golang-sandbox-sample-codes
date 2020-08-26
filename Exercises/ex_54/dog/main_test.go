package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T){
	hy:=2
	dy:=Years(hy)
	if dy != hy*7{
		t.Error("Expected:",hy*7,", Got:",dy)
	}
}

func ExampleYears(){
	dy:=Years(2)
	fmt.Println(dy)
	// Output:
	// 14
}

func BenchmarkYears(b *testing.B){
	for i:=0;i<b.N;i++{
		Years(2)
	}
}

func TestYearsTwo(t *testing.T){
	hy:=2
	dy:=Years(hy)
	if dy != hy*7{
		t.Error("Expected:",hy*7,", Got:",dy)
	}
}

func ExampleYearsTwo(){
	dy:=Years(2)
	fmt.Println(dy)
	// Output:
	// 14
}

func BenchmarkYearsTwo(b *testing.B){
	for i:=0;i<b.N;i++{
		Years(2)
	}
}