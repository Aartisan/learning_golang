package main
import(
	"time"
)
func test(c chan int) {
	time.Sleep(3 * time.Second)
	println("go...")
	c <- 1
}
func main() {
	c := make(chan int)
	go test(c)
	println("hi!")
	data := <- c
	/*<- c   //阻塞等待退出信号*/
	println("over! the return is: ", data )

}
