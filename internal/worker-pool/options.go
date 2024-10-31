package workerpool

type WorkerPool interface {
	Add(delta int)
	Done(subtract int)
	SetWork(work func(n int, chRead <-chan string))
	Count() int
	SetData(<-chan string)
}