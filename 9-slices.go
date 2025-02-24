package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string // dizilerin aksine, dilimler yalnızca içerdikleri elemanlara göre(eleman sayısına göre değil) yazılır.
	// başlatılmamış bir dilim nil'e eşittir ve uzunluğu(length) sıfırdır.
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) // sıfır olmayan uzunlukta bir boş bir dilim oluşturmak için make kullanın.
	// burada uzunluğu(length) 3 olan dizelerden oluşan bir dilim oluşturuyoruz(başlangıçta sıfır değerli).
	// varsayılan olarak yeni bir dilimin kapasitesi uzunluğuna eşittir. dilimin önceden artacağını biliyorsak make'e ek bir
	// parametre olarak açıkça kapasite geçirmek mümkündür.
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a" // dizilerde olduğu gibi ayarlayıp alabiliriz.
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // len, dilimin uzunluğunu beklendiği gibi döndürür.

	s = append(s, "d") // temel işlemlere ek olarak, dilimlerde append'i kullanabiliriz.
	// yeni bir diim değeri alabileceğimiz için, "append" komutundan return değeri kabul etmemiz gerektiğini unutmayın.
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) // dilimler ayrıca kopyalanabilir.
	copy(c, s)                  // burada s ile aynı uzunlukta boş bir c dilimi oluşturuyoruz ve s'den c'ye kopyalıyoruz.
	fmt.Println("cpy:", c)

	l := s[2:5] // dilimler slice[low:high] operatörünü destekler.
	fmt.Println("sl1:", l)

	l = s[:5] // bu dilim s[5]'e kadar (ama 5 hariç) dilimlenir.
	fmt.Println("sl2:", l)

	l = s[2:] // ve bu dilim s[2]'den (2 dahil) dilimlenir.
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // dilim için bir değişkeni tek bir satırda tanımlayıp başlatabiliriz.
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3) // dilimler çok-boyutlu veri yapılarına dönüştürülebilir. iç dilimler'in uzunluğu çok boyutlu
	// dizilerin aksine değişebilir.
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

// Dilimlerin dizilerden farklı tipler olduğunu, ancak fmt.Println tarafından benzer şekilde işlendiğini unutmayın.
