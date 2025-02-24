package main

// go'da multiple return değerleri kullanılır. bu özellik, örneğin bir fonksiyondan
// hem sonuç hem de hata değerlerini döndürmek için sıklıkla kullanılır.

import "fmt"

func vals() (int, int) { // buradaki (int, int) fonksiyonunun 2 int döndürdüğünü gösterir.
	return 3, 7
}

func main() { // burada call'dan gelen 2 farklı return değerini multiple assignment(çoklu atama) ile kullanıyoruz.

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // yalnızca döndürülen değerlerin bir alt kümesini istiyorsak, boş tanımlayıcı "_" kullanın.
	fmt.Println(c)
}
