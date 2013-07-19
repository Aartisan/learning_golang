package main

import (
	"fmt"
)

func Less(i, j interface{}) bool {
	switch i.(type) {
	case int:
		if _, ok := j.(int); ok {
			return i.(int) < j.(int)
		}
	case float32:
		if _, ok := j.(float32); ok {
			return i.(float32) < j.(float32)
		}
	}
	return false
}

func main() {
	var a, b, c int = 3, 4, 0
	if c = a; Less(a, b) {
		c = b
	}

	var x, y, z float32 = 5.5, 3.14, 0.0
	if z = x; Less(x, y) {
		z = y
	}
	fmt.Println(c, z)
}
