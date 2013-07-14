package main

import (
	"fmt"
)

func Map_int(f func(int) int, arr []int) []int {
	x := make([]int, len(arr))
	for i, j := range arr {
		x[i] = f(j)
	}
	return x
}

func Map_string(f func(string) string, arr []string) []string {
	x := make([]string, len(arr))
	for i, j := range arr {
		x[i] = f(j)
	}
	return x
}

func main() {
	arr_int := []int{2, 3, 4, 5, 4, 3, 2}
	f_int := func(i int) int {
		i = i + 1
		return i
	}
	fmt.Printf("%v", Map_int(f_int, arr_int))

	arr_string := []string{"a", "b", "c", "d"}
	f_string := func(x string) string {
		x = "new_" + x
		return x
	}
	fmt.Printf("%v", Map_string(f_string, arr_string))
}
