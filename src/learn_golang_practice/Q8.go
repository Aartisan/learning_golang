package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	//注释下面一句话的目的在于下面的i的作用域出问题了，i的定于位于for的临时变量
	/*fmt.Printf("%v\n", i)*/
}
