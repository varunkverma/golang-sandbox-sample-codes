package main

import (
	"testing"
)

func TestMySum(t *testing.T){

	type test struct{
		data []int
		answer int
	}

	tests:=[]test{
		test{ data:[]int{1,2,3,4}, answer:10 },
		test{ data:[]int{-1,-2,-3,-4}, answer:-10 },
		test{ data:[]int{-1,2,3,-4}, answer:0 },
		test{ data:[]int{}, answer:0 },
	}

	for _,v:=range tests{
		x:=mySum(v.data...)
		if x!=v.answer{
			t.Error("Expected:",v.answer,", got:",x)
		}
	}
	
}