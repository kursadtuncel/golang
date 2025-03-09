package main

import "fmt"

func main() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "Hello World"
	}()
	//blocking
	message := <-myChannel

	fmt.Println(message)
}
