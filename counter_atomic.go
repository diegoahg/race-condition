package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var c int64

func increment() {
	atomic.AddInt64(&c, 1)
}

// Las carreras de datos también pueden ocurrir en variables de tipos primitivos (bool, int, int64, etc.).
func value() int64 { return int64(c) }

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 100; i++ {
				increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Total:", value(), "expected 1000000")
}