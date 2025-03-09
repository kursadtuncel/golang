package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("f1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
	}()

	wg.Wait() //blocked
	fmt.Println("end of the main")
}
