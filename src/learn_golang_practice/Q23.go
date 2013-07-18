/*1. 假设有下面的程序。要注意的是包container/vector曾经是Go的一部分, 但是当内建的 append 出现后,就被移除了。
然而,对于当前的问题这不重要。这个包实现了有 push 和 pop 方法的栈结构。
package main
import "container/vector"
func main() {
	k1 := vector.IntVector{}
	k2 := &vector.IntVector{}
	k3 := new(vector.IntVector)
	k1.Push(2)
	k2.Push(3)
	k3.Push(4)
}
k1,k2 和 k3 的类型是什么?

2. 当前,这个程序可以编译并且运行良好。在不同类型的变量上Push都可以工作。Push 的文档这样描述:
func (p *IntVector) Push(x int) Push 增加 x 到向量的末尾。
那么接受者应当是 *IntVector 类型,为什么上面的代码(Push 语句)可以正确工作?
above (the Push statements) work correct then?*/

package main

import "container/vector"

func main() {
	k1 := vector.IntVector{}
	k2 := &vector.IntVector{}
	k3 := new(vector.IntVector)
	k1.Push(2)
	k2.Push(3)
	k3.Push(4)
}
