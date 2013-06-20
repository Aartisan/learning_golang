package main
import (
	"fmt"
)

type Integer int

func Less (a Integer, b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 1
	if Less(a, 2) {
		fmt.Println(a, "Less 2")
	}
}
