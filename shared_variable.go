package main

import "os"

// ParallelWrite escribe los datos en file1 y file2, devuelve los errores.
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// Este error se comparte con el goroutine principal,
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // La segunda escritura a la misma variable err
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}


func main() {
	ParallelWrite([]byte("Race condition"))
}