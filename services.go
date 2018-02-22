package main

import (
	"fmt"
	"sync"
)


// Variable global desprotegida
// Si hacemos un run del archivo con el flag -race la advertencia de race condition saltara, pero si
// descomentamos las lineas de codigo comentadas, aplicariamos sincronizacion con Mutex y estaria OK.


var service = map[string]string{}
var serviceMu sync.Mutex

func RegisterService(name string, addr string) {
	//serviceMu.Lock()
	//defer serviceMu.Unlock()
	service[name] = addr
}

func LookupService(name string) string {
	//serviceMu.Lock()
	//defer serviceMu.Unlock()
	return service[name]
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// Write
	go func() {
		RegisterService("Payment", "100.123.34.1")
		RegisterService("Order", "100.123.34.2")
		RegisterService("User", "100.123.34.3")
		wg.Done()
	}()

	// Read
	go func() {
		fmt.Println(LookupService("Payment"))
		fmt.Println(LookupService("User"))
		wg.Done()
	}()

	wg.Wait()

}

