package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var c int64

func add() {
	atomic.AddInt64(&c, 1)
}

func value() int64 { return int64(c) }

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 100; i++ {
				add()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Total:", value(), "expected 1000000")
}