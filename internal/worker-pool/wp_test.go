package workerpool

import (
	"sync"
	"testing"
)

type Case struct {
	data       <-chan string
	countWork  int
	work       func(n int, chRead <-chan string)
	resultJobs []string
}

func TestAdd(t *testing.T) {
	datasJobs := Datas()

	resultsJobs := make([][]string, 10, 10)

	datas := make([]<-chan string, 0, 10)

	for _, dataJob := range datasJobs {
		datas = append(datas, DataOfSlice(dataJob))
	}

	cases := []Case{
		Case{
			data:      datas[0],
			countWork: 0,
			work: func(n int, chRead <-chan string) {
				select {
				case str := <-chRead:
					{
						resultsJobs[0] = append(resultsJobs[0], str)
					}
				default:
					{
						return
					}
				}
			},
			resultJobs: datasJobs[0],
		},
	}

	for _, val := range cases {
		wg := NewWP(val.data, val.countWork, val.work)
		wg.Add(10)
	}
}

func (c *Case) Job1(n int, data <-chan string) {
	select {
	case str := <-data:
		{
			c.resultJobs = append(c.resultJobs, str)
		}
	default:
		{
			return
		}
	}
}

func DataOfSlice(strData []string) <-chan string {
	wg := &sync.WaitGroup{}
	result := make(chan string)
	for _, str := range strData {
		wg.Add(1)
		go func(s string) {
			result <- s
			wg.Done()
		}(str)
	}
	wg.Wait()
	close(result)
	return result
}

func Datas() [][]string {
	datasJobs := make([][]string, 0, 10)

	dataJobs1 := []string{
		"a",
		"b",
		"c",
		"d",
		"f",
	}

	dataJobs2 := []string{
		"aa",
		"bb",
		"cc",
		"dd",
		"ff",
	}

	dataJobs3 := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}

	dataJobs4 := []string{
		"How",
		"Where",
		"Data",
		"mmm",
		"12431d",
	}

	datasJobs = append(datasJobs, dataJobs1)
	datasJobs = append(datasJobs, dataJobs2)
	datasJobs = append(datasJobs, dataJobs3)
	datasJobs = append(datasJobs, dataJobs4)

	return datasJobs
}
