package main

import "fmt"

/*A method is a function with an implicit first argument, called a receiver.*/

type Rectangle struct {
	width, height float64
}

/*这是一个func函数*/
func area(r Rectangle) float64 {
	return r.width * r.height
}

/*这是一个method,receiver是一个Rectangle结构体,则表示这method是一个值传递的方法*/
func (r Rectangle) re_area() float64 {
	return r.width * r.height
}

/*这是一个method,receiver是一个Rectangle结构体的指针,则表示这method是一个引用传递的方法*/
func (r *Rectangle) point_re_area() float64 {
	r.width = 10
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	r3 := Rectangle{1, 1}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", r2.re_area())
	fmt.Println("Area of r3 is: ", r3.point_re_area())
	fmt.Println("r3 width = ", r3.width)
}
