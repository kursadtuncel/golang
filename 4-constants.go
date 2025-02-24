package main

import (
	"fmt"
	"math"
)

const s string = "constant" // const ile sabit bir değer tanımlanır.

func main() {
	fmt.Println(s)
	const n = 500000000 // const değeri, var ifadesi bulunan her yerde bulunabilir.
	const d = 3e20 / n  //
	fmt.Println(d)
	fmt.Println(int64(d)) // sayısal sabitin, açıkça tür verilene kadar bir türü yoktur.
	fmt.Println(math.Sin(n))
}
