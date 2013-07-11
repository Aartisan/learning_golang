package main

import (
	"fmt"
)

func average_float64(arr []float64) (avg float64) {
	sum := 0.0
	switch len(arr) {
	case 0:
		avg = 0
	default:
		for _, v := range arr {
			sum = sum + v
		}
		avg = sum / float64(len(arr))
	}
	return
}

func main() {
	arr := []float64{9, 8, 7, 6}
	fmt.Println("average is: ", average_float64(arr))
}
