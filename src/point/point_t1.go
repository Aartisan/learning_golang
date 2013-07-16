package main

import (
	"fmt"
)

func main() {
	var p *int
	fmt.Printf("%v\n",p)
	var i int
	p = &i
	fmt.Printf("%v\n",p)
	*p = 8
	fmt.Printf("%v\n",*p)
	fmt.Printf("%v\n",i)
	//应为golang里面没有指针的运算所以*p++ = (*p)++
	*p++
	fmt.Printf("%v\n",i)
}
