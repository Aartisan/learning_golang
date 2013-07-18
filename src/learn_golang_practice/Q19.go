package main

import (
	"fmt"
)

type e interface{}

func muct(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return "Ray_" + f.(string)
	}
	return f
}

func Map(arr []e, f func(e) e) []e {
	x := make([]e, len(arr))
	for i, j := range arr {
		x[i] = f(j)
	}
	return x
}

func main() {
	int_arr := []e{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}
	string_arr := []e{"a", "b", "c", "d"}

	intf := Map(int_arr, muct)
	stringf := Map(string_arr, muct)
	fmt.Printf("%v\n", intf)
	fmt.Printf("%v\n", stringf)
}
