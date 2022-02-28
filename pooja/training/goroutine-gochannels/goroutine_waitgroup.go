package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	var waitg sync.WaitGroup //wait for multiple goroutines to finish
	for i := 1; i <= 10; i++ {
		waitg.Add(1)
		i := i
		go func() {
			defer waitg.Done()
			task(i)
		}()
	}
	waitg.Wait()
}
