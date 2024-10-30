package main

import (
	"fmt"
	workerpool "testVK/internal/worker-pool"
	"time"
)

func main() {
	getData := make(chan string)
	go func() {
		for {
			getData <- "a"
			getData <- "b"
			getData <- "c"
		}
	}()

	wp := workerpool.NewWP(getData, 600, PrintData)

	wp.Add(500)
	time.Sleep(4 * time.Second)
	wp.Done(300)
	var n int
	fmt.Scan(&n)
}

func PrintData(n int, chRead <-chan string) {
	select {
	case data, ok := <-chRead:
		{
			if ok {
				fmt.Printf("Number Worker:%v data:%s\n", n, data)
			} else {
				return
			}
		}
	default:
		{
			return
		}
	}
}
