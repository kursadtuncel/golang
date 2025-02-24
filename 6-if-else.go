package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // else ifadesi olmadan da if ifadesini kullanabiliriz.
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 { // && ve || gibi mantıksal operatörler koşullarda çok kullanışlıdır.
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 { // Bir ifade koşullu ifadelerden önce gelebilir;
		// bu ifadede bildirilen tüm değişkenler geçerlidir ve sonraki tüm dallarda kullanılabilir.
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Go'da koşulların etrafında parantezlere ihtiyacınız olmadığını, ancak süslü parantezlerin zorunlu olduğunu unutmayın.
