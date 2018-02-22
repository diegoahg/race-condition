package main

import (
	"fmt"
	"sync"
)

// START OMIT
type counter int64
func (c *counter) increment() { *c++ }
func (c *counter) value() int64 { return int64(*c) }

func main() {
	var myCounter counter
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 100; i++ {
				// Session critica
				myCounter.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Total:", myCounter.value(), "expected 1000000")
	// END OMIT
}