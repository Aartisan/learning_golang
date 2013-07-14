package main

import (
	"fmt"
)

func max(arr []int) int {
	max := 0
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func min(arr []int) int {
	var min int
	min  = arr[0]
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return min
}

func main() {
	arr := []int{5,6,71,23,46,9,234,64}
	fmt.Printf("In arr max is %d\n", max(arr))
	fmt.Printf("In arr min is %d\n", min(arr))
}
