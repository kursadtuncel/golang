package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { // Aynı case ifadesinde birden fazla ifadeyi ayırmak için virgül kullanabiliriz.
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now() // "switch" ifadesi olmadan if/else mantığını ifade etmenin alternatif bir yoludur.
	//Burada ayrıca case ifadelerinin sabit olmayan ifadeler olabileceğini de gösteriyoruz.
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		// switch, değerler yerine türleri karşılaştırır. bunu  bir arayüz değerinin türünü öğrenmek için kullanabiliriz.
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
			// bu örnekte t değişkeni, kendisine karşılık gelen türe sahip olacaktır.
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
