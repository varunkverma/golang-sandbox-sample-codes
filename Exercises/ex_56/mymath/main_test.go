package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T){
	type test struct{
		data []int
		answer float64
	}

	tests:=[]test{
		test{ data:[]int{1, 4, 6, 8, 100}, answer:6 },
		test{ data:[]int{9000, 4, 10, 8, 6, 12}, answer:9},
		test{ data:[]int{23, 744, 140, 200}, answer:170 },
	}
	var actual float64
	for _,tester:=range tests{
		actual=CenteredAvg(tester.data)
		if actual!=tester.answer{
			t.Error("Expected:",tester.answer,", Got:",actual)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B){
	for i:=0;i<b.N;i++{
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}

func ExampleCenteredAvg(){
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}

