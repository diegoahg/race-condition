package main

import (
	"fmt"
)

// No te olvides de capturar las variables de interacción.

type file struct { name string }
func (f *file) close() { fmt.Println("Close bill: ", f.name) }

func main() {
	files := make([]*file, 0, 1000)

	for i := 0; i < cap(files); i++ {

		files = append(files, &file{
			name: fmt.Sprintf("Bill %v", i+1),
		})
	}

	for _, f := range files {
		go func() {
			f.close()
		}()
	}
}