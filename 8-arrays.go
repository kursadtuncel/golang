package main

import "fmt"

func main() {
	var a [5]int
	// Burada tam olarak 5 tam sayı tutacak bir a dizisi oluşturuyoruz. Elemanların türü ve uzunluğu dizinin türünün bir parçasıdır.
	// Varsayılan olarak bir dizi sıfır değerlidir, bu da tam sayılar için 0'lar anlamına gelir.
	fmt.Println("emp:", a)

	a[4] = 100
	// array[index] = value sözdizimini kullanarak index'e bir değer atayabilir ve array(index) ile bir değer alabiliriz.
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("get:", len(a)) // burada yerleşik len, dizinin uzunluğunu döndürür.

	b := [5]int{1, 2, 3, 4, 5} // bir diziyi tek satırda bildirmek ve başlatmak için bu sözdizimini kullanıyoruz.
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5} // ayrıca "..." ile derleyicinin eleman sayısını saymasını da sağlayabiliriz.
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("dcl:", b) // eğer index'i ":" ile belirtirsek, aradaki elemanlar sıfırlanacaktır.

	var twoD [2][3]int // dizi tipleri tek boyutludur, ancak çok-boyutlu veri yapıları oluşturmak için tipleri birleştirebiliriz.
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{{1, 2, 3}, {4, 5, 6}} // çok-boyutlu dizileri aynı anda oluşturabilir ve başlatabiliriz.

	fmt.Println("2d: ", twoD)
}

// fmt.Println ile yazdırıldığında dizilerin [v1 v2 v3] şeklinde göründüğüne dikkat edin.
