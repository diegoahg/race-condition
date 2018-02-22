package main

import (
	"fmt"
	"sync"
)

// START OMIT
var mutex sync.Mutex
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
				mutex.Lock()
				myCounter.increment()
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Total:", myCounter.value(), "expected 1000000")
	// END OMIT
}