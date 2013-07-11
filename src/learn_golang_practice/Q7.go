package main

import (
	"fmt"
)

func sort(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

func main() {
	i, j := sort(2, 7)
	fmt.Printf("f(2,7) return is %d, %d\n", i, j)
	i, j = sort(7, 2)
	fmt.Printf("f(7,2) return is %d, %d", i, j)
}
