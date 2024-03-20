package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var chooseTask int
	var numReceivers int

	flag.IntVar(&chooseTask, "chooseTask", 2, "choose a task")
	flag.IntVar(&numReceivers, "numReceivers", 1, "number of receivers")
	flag.Parse()

	switch chooseTask {
	case 1:
		task1()
	case 2:
		task2(numReceivers)
	default:
		fmt.Println("Неправильный ввод")
	}
}

func task1() {
	ch := make(chan int)
	done := make(chan struct{})
	defer close(done)

	go receiverFirstTask(ch, done)

	const numGoroutines = 100
	const maxParallel = 10
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxParallel)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			for j := 0; j < 10; j++ {
				value := workerID + j
				ch <- value

				fmt.Printf("Горутина %d: Кладет %d в канал\n", workerID, value)

				randSleep := time.Duration(rand.Intn(1000)+10) * time.Millisecond
				time.Sleep(randSleep)

				fmt.Printf("Горутина %d завершила свою работу!\n", workerID)
			}
		}(i)
	}

	wg.Wait()
	close(ch)
}

func receiverFirstTask(ch <-chan int, done <-chan struct{}) {
	for {
		select {
		case value := <-ch:
			fmt.Println("Ресивер получил", value)
		case <-done:
			return
		}
	}
}

func task2(numReceivers int) {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(numReceivers)

	go publisherSecondTask(ch)

	for i := 0; i < numReceivers; i++ {
		go func() {
			defer wg.Done()
			receiverSecondTask(ch)
		}()
	}

	wg.Wait()
}

func publisherSecondTask(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i

		fmt.Printf("Паблишер отправил сообщение: %d\n", i)

		time.Sleep(time.Millisecond * 500)
	}

	close(ch)
}

func receiverSecondTask(ch <-chan int) {
	for {
		value, ok := <-ch

		if !ok {
			return
		}

		fmt.Println("Ресивер получил", value)
	}
}
