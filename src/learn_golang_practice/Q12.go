package main

import(
	"fmt"
)

func Map(f func(int) int,arr []int) []int {
	x := make([]int, len(arr))
	for i, j := range arr {
		x[i] = f(j)
	}
	return x
}

func main() {
	arr := []int{2,3,4,5,4,3,2}
	f := func(i int) int {
		i = i + 1
		return i
	}
	fmt.Printf("%v", Map(f, arr))
}
