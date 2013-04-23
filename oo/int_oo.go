package main
import (
	"fmt"
)

type integer int

func (a integer) less (b integer) bool {
	return a < b
}

func main() {
	var a integer = 1
	if a.less(2) {
		fmt.Println(a, "less 2")
	}
}
