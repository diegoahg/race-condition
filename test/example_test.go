package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// START OMIT
// https://blog.golang.org/race-detector
func TestTimer(t *testing.T) {
	start := time.Now()
	var timer *time.Timer
	timer = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		timer.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

// END OMIT
func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
