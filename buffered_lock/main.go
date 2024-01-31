package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(lock chan bool, id int, done chan struct{}) {
	fmt.Printf("%d started\n", id)

	lock <- true

	fmt.Printf("%d is acquired the lock\n", id)
	duration := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("%d is released the lock\n", id)

	done <- struct{}{}

	<-lock
}

func main() {
	//fmt.Println("Buffered lock")
	lock := make(chan bool, 1)

	done := make(chan struct{})

	for i := 0; i < 5; i++ {
		go worker(lock, i, done)
	}

	for i := 0; i < 5; i++ {
		<-done
	}

}
