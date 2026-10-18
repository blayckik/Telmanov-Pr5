package main

import (
	"fmt"
	"sync"
)

func main() {
	urgent := make(chan int, 10)
	normal := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		normal <- i
		urgent <- i
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			select {
			case x := <-urgent:
				fmt.Println("Срочная задача", x)
			default:
				x := <-normal
				fmt.Println("Обычная задача", x)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Все задачи выполнены")
}