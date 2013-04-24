package main
import (
	"fmt"
)


func main() {
	var a =  [3]int{1,2,3}
	var b = &a
	fmt.Println(b[1])
}
