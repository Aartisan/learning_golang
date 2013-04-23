package main
import (
	"fmt"
)

type Integer int

func (a *Integer) Add_point (b Integer) {
	*a += b
}

func (a Integer) Add_no_point (b Integer) {
	a += b
}

func main() {
	var a Integer = 1
	var b Integer = 1
	a.Add_point(2)
	b.Add_no_point(2)
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
}
