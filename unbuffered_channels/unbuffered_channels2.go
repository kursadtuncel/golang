package main

import "fmt"

func main() {
	myChannel := make(chan string)
	done := make(chan bool)
	go func() {
		message := <-myChannel
		fmt.Println(message)
		done <- true
	}()

	go func() {
		myChannel <- "Hello World"
	}()
	<-done
	fmt.Println("End of the main")

}
