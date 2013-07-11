package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[index => " + strconv.Itoa(i) + "; value => " + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {
	s := new(stack)
	/*var s stack*/
	s.push(3)
	s.push(4)
	s.push(2)
	i := s.pop()
	fmt.Printf("pop value is %v\n", i)
	s.push(1)
	fmt.Printf("stack %v\n", s)
}
