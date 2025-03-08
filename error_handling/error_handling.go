package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		fileData, err := os.ReadFile("sample.txt")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fileData)
		}
	*/
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("payda sıfır olamaz")
	}
	return x / y, nil
}
