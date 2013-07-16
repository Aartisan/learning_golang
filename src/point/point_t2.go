//new 分配;make 初始化 
//上面的两段可以简单总结为:
//		1. new(T) 返回 *T 指向一个零值 T 
//		2. make(T) 返回初始化后的 T
//当然 make 仅适用于 slice,map 和 channel。

