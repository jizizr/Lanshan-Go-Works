package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	current := 0

	printEven := func() {
		for {
			mu.Lock()
			for current%2 != 0 {
				cond.Wait()
			}
			if current >= 100 {
				mu.Unlock()
				return
			}
			fmt.Println(current)
			current++
			cond.Signal()
			mu.Unlock()
		}
	}

	printOdd := func() {
		for {
			mu.Lock()
			for current%2 == 0 {
				cond.Wait()
			}
			if current >= 100 {
				mu.Unlock()
				return
			}
			fmt.Println(current)
			current++
			cond.Signal()
			mu.Unlock()
		}
	}

	go printEven()
	go printOdd()
	for current < 100 {
		time.Sleep(100 * time.Millisecond)
	}
}
