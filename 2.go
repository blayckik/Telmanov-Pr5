package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-done:
			fmt.Println("Горутина остановлена")
			return
		case <-time.After(time.Second):
			fmt.Println("Работаю...")
		}
	}
}

func main() {
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(done, &wg)

	time.Sleep(3 * time.Second)
	fmt.Println("Отправляю сигнал остановки")
	done <- true

	wg.Wait()
}