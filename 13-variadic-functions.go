package main

// değişkenli (variadic) fonksiyonlar herhangi sayıda ardışık argümanla çağrılabilir.
// örneğin fmt.Println yaygın bir değişkenli fonksiyondur.

import "fmt"

func sum(nums ...int) { // argüman olarak rastgele sayıda int alacak bir fonksiyon
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums { // fonksiyon içerisinde num tipi []int'e eşdeğerdir.
		// len(nums)'ı çağırabilir, range ile üzerinde iterasyon yapabiliriz.
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2) // değişkenli fonksiyonlar, her zamanki gibi, ayrı argümanlarla çağırılabilir.
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4} // bir dilimde birden fazla argümanınız varsa, bunları func(slice...) kullanarak
	// değişkenli bir fonksiyona bu şekilde uygulayabiliriz:
	sum(nums...)
}
