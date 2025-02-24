package main

// go, iç içe fonksiyonlar oluşturabilen anonim fonksiyonları destekler.

import "fmt"

func intSeq() func() int { // bu intSeq fonksiyonu, intSeq'in gövdesinde anonim olarak tanımladığımız
	// başka bir fonksiyonu döndürür. döndürülen fonksiyon, bir kapanış oluşturmak için i değişkeni üzerinde kapanır.
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
