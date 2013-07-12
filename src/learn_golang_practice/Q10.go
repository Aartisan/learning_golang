package main

import (
	"fmt"
)

func special_print(arr ...int) {
	for _, d := range arr {
		fmt.Printf("%d\n", d)
	}
}

func special_print_new(arr ...int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\n", arr[i])
	}
}

func main() {
	special_print(1, 2, 2, 3, 4, 5, 1, 3, 1, 3)
	special_print_new(1, 2, 2, 3, 4, 5, 1, 3, 1, 3)
}
