package main

import (
	"fmt"
	workerpool "testVK/internal/worker-pool"
	"time"
)

func main() {
	getData := make(chan string)
	go func() {
		getData <- "a"
		getData <- "b"
		getData <- "c"

		close(getData)
	}()

	wp := workerpool.NewWP(getData, 600, PrintData)

	wp.Add(500)
	time.Sleep(4 * time.Second)
	wp.Done(5)
	time.Sleep(4 * time.Second)
	wp.SetWork(AnotherJob)
	time.Sleep(4 * time.Second)

	data := make(chan string)
	go func() {
		for {
			data <- "ff"
			data <- "kk"
			data <- "ll"
		}
	}()

	wp.SetData(data)
	var n int
	fmt.Scan(&n)
	j(wp)
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

func AnotherJob(n int, chRead <-chan string) {
	select {
	case data, ok := <-chRead:
		{
			if ok {
				fmt.Printf("Another job: data:%s\n", data)
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

func j(wor workerpool.WorkerPool) {

}