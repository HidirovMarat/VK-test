# VK test

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](http://forthebadge.com/images/badges/built-with-love.svg)](http://forthebadge.com)

Реализовать примитивный worker-pool с возможностью динамически добавлять и удалять воркеры. Входные данные (строки) поступают в канал, воркеры их обрабатывают (например, выводят на экран номер воркера и сами данные). Задание на базовые знания каналов и горутин.

Используемые технологии:
- golang/mock, testify (для тестирования)

Сервис разработан с использованием Clean Architecture, что обеспечивает простоту расширения его функциональности и удобство тестирования.

# Запуск
Через run - `go run cmd/app/main.go` или build - `go build -o <your desired name>`.  

## Examples
Пример работы с пакетом workerpool
```go
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
```
Нам нужен канал откуда будем читать данные и необходимые значения для конструктора(NewWP).
# Решения <a name="decisions"></a>
В ходе разработки был сомнения по тем или иным вопросам, которые были решены следующим образом:
- Все условия тестового задания были выполнены, то есть добавление и удаление worker динамически реализованы с помощью методов 'Add' и 'Done' соответственно.
- Было решено добавить возможность динамической смены job для worker с помощью метода 'SetWork'.
- Также добавлено возможность смены канала данных динамически.

# Тесты и Бенчмарки

