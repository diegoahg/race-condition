package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	start := time.Now()
	// Se agrega un channel para implementar una aprox de 'Pasaje de Mensajes'
	reset := make(chan bool)
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		reset <- true
	})
	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
