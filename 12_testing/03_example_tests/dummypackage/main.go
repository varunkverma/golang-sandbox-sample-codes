// Package dummypackage is just to implement Example Tests
package dummypackage

// Sum takes an integer variadic parameter i.e., infinite number of integers and return their integral sum 
func Sum(nums... int) int {
	s:=0
	for _,v:=range nums{
		s+=v
	}
	return s
}