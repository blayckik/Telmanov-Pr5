package main

import (
	"fmt"
	"time"
)

func request(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "ответ сервера"
}

func main() {
	ch := make(chan string)

	go request(ch)

	select {
	case res := <-ch:
		fmt.Println("Получен:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Таймаут")
	}
}