package main

import "time"

func main() {
	c := make(chan int)
	o := make(chan bool)
	tick := time.Tick(1 * time.Second)
	timeout := time.After(5 * time.Second)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-tick:
				println("tick")
			case <-timeout:
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<- o
}
