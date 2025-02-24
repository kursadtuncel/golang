package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // boş bir map oluşturmak için make komutunu kullanıyoruz.
	// make(map[key-type]val-type)

	m["k1"] = 7 // tipik name[key] = value sözdizimini kullanarak key/value(anahtar/değer) çiftlerini ayarlayın.
	m["k2"] = 13

	fmt.Println("map:", m) // fmt.Println ile map'i yazdırmak, map'in tüm anahtar/değer çiftlerini gösterir.

	v1 := m["k1"] // name[key] ile bir değer alıyoruz.
	fmt.Println("v1:", v1)

	v3 := m["k3"] // key mevcut değilse, sıfır değeri döndürülür.
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m)) // len, map'te çağırıldığında key/value çiftlerinin sayısını döndürür.

	delete(m, "k2") // delete ile map'ten key/value çiftlerini silebiliriz.
	fmt.Println("map:", m)
	clear(m) // map'in tüm key/value çiftlerini silmek için clear kullanabiliriz.
	fmt.Println("map", m)

	_, prs := m["k2"] // Bir haritadan değer alırken isteğe bağlı ikinci dönüş değeri, anahtarın haritada mevcut olup olmadığını gösterir.
	// Bu, eksik anahtarlar ile 0 veya "" gibi sıfır değerli anahtarlar arasındaki belirsizliği gidermek için kullanılabilir. Burada değerin kendisine ihtiyacımız yoktu, bu yüzden boş tanımlayıcı _ ile onu görmezden geldik.
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // bu sözdizimini kullanarak aynı satırda yeni bir map'i tanımlayıp başlatabiliriz.
	fmt.Println("map", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

// map'in fmt.Prinln ile yazdırıldığında map[k:v k:v] biçminde çıktı verdiğine dikkat edin.
