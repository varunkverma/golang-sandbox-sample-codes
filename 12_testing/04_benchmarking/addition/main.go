// Package addition does addition
package addition

// AddAll takes unlimited numbers of integer types and returns their integral sum 
func AddAll(nums... int) int {
	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	return sum
}