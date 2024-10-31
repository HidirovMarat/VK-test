package workerpool

import (
	"log"
	"sync"
	"time"
)

type WP struct {
	data    <-chan string
	signals []chan struct{}
	work    func(n int, chRead <-chan string)
	wg      sync.WaitGroup
}

func NewWP(data <-chan string, countWork int, work func(n int, chRead <-chan string)) *WP {
	return &WP{
		data:    data,
		signals: make([]chan struct{}, 0, countWork),
		work:    work,
	}
}

func (wp *WP) Add(delta int) {
	if delta < 1 {
		return
	}

	log.Printf("Add workers %v\n", delta)
	for i := 0; i < delta; i++ {
		signal := make(chan struct{})
		go func(n int, signal <-chan struct{}) {
			for {
				select {
				case <-signal:
					{
						log.Printf("==Close Worker %v==\n", n)
						return
					}
				default:
					{
						time.Sleep(1 * time.Second)
						wp.work(n, wp.data)
					}
				}
			}
		}(len(wp.signals)+1, signal)

		wp.signals = append(wp.signals, signal)
	}
}

func (wp *WP) Done(subtract int) {
	if subtract < 1 {
		return
	}

	if len(wp.signals) < subtract {
		subtract = len(wp.signals)
	}

	log.Printf("Stop workers %v\n", subtract)
	for i := len(wp.signals) - 1; i >= len(wp.signals)-subtract; i-- {
		wp.wg.Add(1)
		go func() {
			wp.signals[i] <- struct{}{}
			wp.wg.Done()
		}()
	}
	wp.wg.Wait()

	wp.signals = wp.signals[:len(wp.signals)-subtract]
}

func (wp *WP) SetWork(work func(n int, chRead <-chan string)) {
	wp.work = work
}

func (wp *WP) Count() int {
	return len(wp.signals)
}

func (wp *WP) SetData(data <-chan string) {
	wp.data = data
}
