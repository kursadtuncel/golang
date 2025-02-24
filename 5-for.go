package main

import "fmt"

func main() {

	i := 1 // tek bir koşula sahip en temel tür.
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { // klasik initial/condition/loop döngüsü.
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
		// "bunu n kere yap" demenin farklı bir yolu olarak tam sayı üzerinde aralık oluşturma.
	}

	for { // for bir koşul olmadan kullanıldığında, döngüden çıkana kadar veya fonksiyon dönene kadar sürekli döngüye girer.
		fmt.Println("loop")
		break
	}

	for n := range 6 { //döngünün bir sonraki iterasyonuna da devam edebiliriz.
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
