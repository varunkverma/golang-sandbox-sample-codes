package addition

import (
	"testing"
	"fmt"
)

func TestAddAll(t *testing.T){
	s:=AddAll(1,2,3,4)
	if s != 10 {
		t.Error("Got:",s,", Expected:",10)
	}
}

func ExampleAddAll(){
	fmt.Println(AddAll(1,2,3,4))
	// Output:
	// 10
}

func BenchmarkAddAll(b *testing.B){
	for i:=0;i<b.N;i++{
		AddAll(1,2,3,4)
	}
}