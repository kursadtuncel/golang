package main

import "fmt"

func main() {
	/*
		var numbers = []int{1, 2, 3, 4}

			for index := 0; index < len(numbers); index++ {
				fmt.Println(numbers[index])
			}


		//üstteki örnek yerine foreach kullanarak:

		for _, value := range numbers {
			fmt.Println(value)
		}
	*/

	/*
		var language = "Golang"

		for _, character := range language {
			fmt.Println(character)
		}
	*/

	var names = map[string]int{
		"Ali":    10,
		"Ahmet":  20,
		"Mehmet": 30,
	}

	for key, value := range names {
		fmt.Println(key, value)
	}

}
