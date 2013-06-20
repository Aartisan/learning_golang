package main

import (
	"fmt"
	"flag"
	"io"
	"bufio"
	"os"
	"strconv"
	"time"
)

import "algorithm/bubblesort"
import "algorithm/qsoralgorithm/qsortt"

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile String)(Values []int, err error){
	if err != nil {
		fmt.Printf("Failed open infile")
		return
	}

	defer file.close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix , err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line ,seems unexcepted")
			return
		}

		str := string(line) //将字符数组转换成字符串
		values, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
}

func writeValues(values []int, outfile string) error {
	file, error := os.Create(outfile)
	if err != nil {
		fmt.Printf("Failed create outfile")
		return
	}

	defer file.Close()

	for _, values := range values{
		str := strconv.Itoa(values)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile = ", *infile,  "outfile =", *outfile, "algorithm =", *algorithm)
		file, err := os.File(infile)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(vales)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unkonwn or unsupported")
		}
		t2 := time.Now()
		fmt.Println("The sorting process cost", t2.Sub(t1), "to complete"
		writeValues(values, *outfile)
	}else{
		fmt.Println(err)
	}
}
