package workerpool

import "testing"

type Case struct {
	data <-chan string
	countWork int
	work func(n int, chRead <-chan string)
	expectedCountWorker int
	resultJobs []string
}

func TestAdd(t *testing.T) {
	//wp := NewWP()
}