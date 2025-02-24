package main

import "fmt"

func plus(a int, b int) int { // 2 int değeri alan ve bunların toplamını int olarak döndüren bir fonksion.
	return a + b // go'da illa ki return belirlemek gerektirir, yani son ifadenin değerini otomatik olarak döndürmez.
}

func plusPlus(a, b, c int) int { // aynı tipten birden fazla ardışık parametremiz olduğunda,
	// tipi tanımlayan son parametreye kadar benzer türdeki parametreler için tip adını atlayabiliriz.
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
