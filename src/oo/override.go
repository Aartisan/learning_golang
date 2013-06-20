package main

type A struct {
	x int
}

type B struct {
	A
	x int
}

func (this *A) Test() {
	println("A:", this.x)
}

func (this *B) Test() {
	println("B:", this.x)
}

func main() {
	o := B{x: 123, A: A{321}}
	o.Test()
	o.A.Test()
}

