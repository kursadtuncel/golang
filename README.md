
# Neden Go?

Go, büyük ve için ölçeklenebilir web uygulamaları inşa edebilmek için geliştirildi.
Bu yüzden Golang geliştirilirken mümkün olduğunca standart hale getirildi, gofmt aracı ve dilin katı kullanım sebebi de geliştirici tabanında iki fraksiyon olmaması içindir.
<br>

# Başlıklar:

*[Hello World](#hello-world)


<br>

# Hello World

Elbette, Hello World ile Go'ya giriş yapıyoruz. 

```golang
package main

import "fmt"

func main() {
    
fmt.Printf("Hello, world!")

}
```
Çıktı:
```golang
Hello, world!
```
___
<br>

# Değişkenler (Variables)

Temel şekilde **var** ile değişken atıyoruz. Go'da değişken türünü değişken adından sonra belirttiğimize dikkat edin.

```golang
var variableName type

//Birden fazla değişken tanımlamak:
var vname1, vname2, vname3 type

//Başlangıç değeri olan bir değişken tanımlamak:
var variableName type = value

//Başlangıç ​​değerleri ile birden fazla değişken tanımlamak:
var vname1, vname2, vname3 type = v1, v2, v3

// "type=" olmadan da tanımlanabilir:
var vname1, vname2, vname3 = v1, v2, v3

// ':=' kullanarak hızlı atama yapabiliriz, bu yalnızca fonksiyon içinde geçerlidir.
vname1, vname2, vname3 := v1, v2, v3

// Fonksiyon dışında kullanmaya çalışırsak derleme hatası alırız.
// Bu nedenle global değişkenleri tanımlamak için kullanmalıyız.

````
**_(blank)** özel bir değişkendir. Buna verilen herhangi bir değer göz ardı edilecektir. Örneğin, b ye 35 veririz ve 34 ü atarız. Bu örnek sadece nasıl çalıştığını gösterir. Burada işe yaramaz görünüyor çünkü genellikle fonksiyon dönüş değerlerini aldığımızda bu sembolü kullanırız.

```golang
_, b := 34, 35
```

# Constants (Sabitler)

Derleme sırasında belirlenen ve runtime(çalışma) sürecinde değiştiremeyeceğimiz değerlerdir.
Go'da sabit türleri olarak sayı, boolean veya dize kullanabiliriz.

```golang
const constantName = value
// gerektiğinde sabit türlerini atlayabiliriz
const Pi float32= 3.1415926

```

# Elementary Types (Veri Tipleri)
## Boolean
Go'da bir değişkeni boolean olarak tanımalamak için **bool** kullanırız, değer yalnızca **true** veya **false** olabilir ve **false** varsayılan değer olacaktır. 
**Dipnot:(Değişkenlerin türünü sayı ve boolean arasında dönüştüremezsiniz.)**

```golang
var isActive bool //global değer
var enabled, disabled = true, false //değişken türünü atla

func test(){
    var available bool //local değer
    valid := false // değişkenin kısa ifadesi
    available := true // değişkene değer ata
}

```

## Numerical Types (Sayısal Veri Tipleri)
integer veri tipleri hem işaretli hem de işaretsiz tam sayı tiplerini içerir.
Go'da aynı anda hem **int** hem de **uint** vardır, aynı boyuta sahiptirler, ancak
belirli boyut işletim sisteminize bağlıdır.
32-bit işletim sistemlerinde 32bit, 64bit işletim sistemlerinde 64bit kullanılır.
Go'da ayrıca rune, int8, int16, int32, int64, byte, uint8, uint16, uint32, uint64 gibi belirli boyutlara sahip tipler de vardır. rune'un int32'nin, byte'ın ise uint8'in takma adı(alias) olduğunu unutmayın.

Bilmeniz gereken en önemli şey, bu değerler arasında değer atayamayacağınızdır, böyle bir işlem derleme hatasına yol açacaktır.

````
var a int8
var b int32
c := a + b
````
int32, int8'den daha fazla bir boyuta sahip olmasına ve int ile aynı türe sahip olmasına rağmen, aralarında değer atayamazsınız. (c burada int türü olarak doğrulanacaktır)

Go karmaşık sayıları da destekler. complex128 (64 bit gerçek ve 64 bit sanal parçayla) varsayılan türdür, daha küçük bir türe ihtiyacınız varsa complex64 (32 bit gerçek ve 32 bit sanal parçayla) adında bir tür daha vardır.

````
var c complex64 = 5+5i
//output: (5+5i)
fmt.Printf("Value is: %v", c)
````

## Strings (Dizeler)
Dizeler çift tırnak işareti **""** veya ters tırnak işareti **``** ile tanımlanır.

```golang
var frenchHello string //string tanımlamanın temel hali
var emptyString string = "" //boş string tanımlamak
funct test() {
    no, yes, maybe := "no", "yes", "maybe"
    japaneseHello := "ohaiou"
    frenchHello := "Bonjour"
}
```

String değerlerini index'e göre değiştirmek imkansızdır. Aşağıdaki kodu derlerseniz, hata alırsınız:
```golang
var s string = "Hello"
s[0] = 'c'

```
Ya bir string değerini gerçekten değiştirmek istiyorsak? Aşağıdaki kodu deneyin:
```golang
s := "hello"
c := []byte(s) // string'i byte tipiyle değiştiriyoruz
c[0] = 'c'
s2 := string(c) // string tipine geri dönüştürüyoruz
fmt.Printf("%s\n", s2)
```

# Error Types (Hata Türleri)
Go'da hata türleriyle başa çıkmak için bir adet hata türü vardır. Ayrıca **errors** adında bir paket mevcuttur.

```golang
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
fmt.Print(err)
}
```

# array, slice, map
## array
Arrays(diziler) şöyle tanımayabiliriz:
```golang
var arr [n]type
```
[n]type'da n array'ın boyutudur, type ise array türüdür. Diğer dillerde olduğu gibi array d eğerlerini almak için [] kullanırız.

```golang
var arr [10]int
arr[0] = 42 // array 0 tabanlıdır
arr[1] = 13 // öğeye değer atama
fmt.Printf("The first element is %d\n", arr[0])
// öğe değerini çağırdık, 42 sayısını döndürecektir
fmt.Printf("The last element is %d\n", arr[9])
// 10. öğenin varsayılan değerini döndürecektir, ki bu sıfırdır.
```

Çünkü boyut(length) dizi türünün bir parçası olduğu için [3]int ve [4]int farklı türlerdir, bu nedenle dizilerin boyutunu değiştiremeyiz. Diziler argüman olarak kullanıldığında, fonksiyonlar referanslar yerine kopyalarını alır. Referansları kullanmak istiyorsak Slice kullanmak isteyebilirsiniz. Bu konuya ilerde değineceğiz.

## slice
Birçok senaryoda dizi(array) iyi bir tercih olmaz. Özellikle diziyi tanımlarken boyutunun ne olacağını bilmediğimizde. Bu nedenle daha dinamik bir diziye ihtiyacımız vardır. Buna slice denir.
Slice gerçekten de bir dinamik dizi değildir. Bir referans türüdür. Slice diziye benzer olan ancak boyutları değişken olan temel bir diziyi işaret eder.

```golang
var fslice []int // tıpkı dizi tanımlama gibi, ancak length'i hariç tutuyoruz.
```

Bir slice tanımlayalım:
```golang
slice := []byte {'a', 'b', 'c', 'd'} 
// slice var olan slice'ları veya dizileri tekrar tanımlayabilir. slice, array[i:j] kullanır,
// burada i başlangıç indeksi ve j bitiş indeksidir.
```

```golang
var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i','j'}
var a,b []byte
a = ar[2:5]
b = ar[3:5]
```
Slice ve Array tanımlarken aradaki farklara dikkat edin. Go'nun boyut hesaplamasını sağlamak için [...] kullanırız ancak sadece slice tanımlamak için [] kullanırız.

### Bazı kullanışlı slice işlevleri:

```golang
var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i','j'}
// 2 tane slice tanımlıyoruz:
var aSlice, bSlice []type

aSlice = array[:3] // aSlice = array[0:3]'e eşittir. aSlice a,b,c öğelerine sahiptir
aSlice = array[5:] // aSlice = array[5:10] aSlice f,g,h,i,j öğelerine sahiptir.
aSlice = array[:] // aSlice = array[0:10] aSlice tüm öğelere sahiptir.

//slice from slice

aSlice = array[3:7] //aSlice d,e,f,g öğelerine sahiptir. len=4, cap=7 dir.
bSlice = aSlice[1:3] // bSlice, aSlice[1], aSlice[2]'yi içerir, yani e,f öğelerine sahiptir.
bSlice = aSlice[:3] // bSlice, aSlice[0], aSlice[1], aSlice[2]'yi içerir, yani d,e,f öğelerine sahiptir.
bSlice = aSlice[0:5] // slice, cap aralığında genişletilebilir, yani şimdi bSlice d,e,f,g,h öğelerine sahiptir.
bSlice = aSlice[:] // bSlice, aSlice ile aynı öğelere sahiptir, yani d,e,f,g

```

Slice bir referans türüdür. Bu nedenle herhangi bir değişiklik halinde aynı slice'a veya diziye işaret eden diğer değişkenler de etkilenir. Yukarıdaki örnekte olduğu gibi, aSlice'daki bir öğrenin değerini değiştirirseniz, bSlice'da değişecektir.

Slice tanım gereği yapı(struct) gibidir ve 3 bölümden oluşur:
- Slice'ın nerede olduğunu belirten bir işaretçi,
- Slice'ın boyutu,
- Kapasitesi, slice'ın başlangıç indeksinden bitiş indeksine kadar olan boyutu.

```
Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i','j'}
Slice_a := Array_a[2:5]
```

Slice için bazı yerleşik fonksiyonlar vardır:
- len slice'ın boyutunu alır.
- cap slice'ın maksimum boyutunu alır.
- append slice'a bir veya daha fazla öğe ekleyebilir, ve slice öğesini döndürür.
- copy öğeleri bir slice'dan diğerine kopyalar ve kopyalanan öğelerin sayısını döndürür.

**Dipnot: append, slice'ın işaret ettiği diziyi değiştirecek ve aynı diziye işaret eden diğer slice'ları da etkileyecektir. Ayrıca slice için yeterli boyut yoksa ((cap-len) == 0), append bu slice için yeni bir dizi döndürür. Bu olduğunda, eski diziye işaret eden diğer slice'lar etkilenmeyecektir.**

## map
map içerisindekei **get** ve **set** değerleri slice'a benzer, fakat slice içindeki indeks yalnızca integer olabilirken, map bundan daha fazlasını(string vb.) kullanabilir.

````golang
var numbers map[string] int
// map tanımlamanın farklı bir yolu:
numbers := make(map[string]int)
numbers["one"] = 1
numbers["ten"] = 10
numbers["three"] = 3
fmt.Println("The third number is: ", numbers["three"])
````
### map kullanırken:
- Map'ler düzensizdir, her yazdırdığımızda farklı sonuçlar elde ederiz. Değerleri indekse göre almak imkansızdır, bu nedenle key kullanmalıyız.
- Map sabit bir boyuta sahip değildir, yani slice gibi bir referans türüdür.
- 'len' map için de çalışır, o map'in kaç tane s anahtarına sahip olduğunu döndürür.
- Değerleri map üzerinden değiştirmek çok kolaydır. Basitçe **numbers["one"]=11** kullanarak anahtarın değerini 11 olarak değiştirebiliriz.

Map değerlerini başlatmak için key:val kullanabiliriz. Map'te anahtarın mevcut olup olmadığını kontrol edebilmek için bazı yöntemleri vardır.

```golang
// map'teki bir öğeyi silmek için "delete" kullanın.
rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2}
// map'in iki return değeri vardır, ikinci return değeri için key yoksa, false döndürür. aksi halde true döndürür.
csharpRating, ok := rating["C#"]
if ok {
    fmt.Println("C# is in the map and its rating is: ", csharpRating)
} else {
    fmt.Println("We have no rating associated with C# in the map")
}

delete(rating, "C") // c içeren key'i siler
```
Yukarıda belirttiğimiz gibi, map bir referans türüdür ve iki map aynı veriye işaret ediyorsa herhangi bir değişiklik ikisini de etkileyecektir.

```golang
m := make(map[string]string)
m["Hello"] = "Bonjour"
m1 := m
m1["Hello"] = "Salut" // m["Hello"] anahtarı 'salut' olarak değişti
```

# make, new
make, map, slice ve channel gibi yerleşik(built-in) modeller için bellek ayırma işlemini gerçekleştirirken, new; türlerin bellek ayırma işlemlerini gerçekleştirir.

new(T) sıfır değerini ayırır, T tipinin belleğine işaret eden bir işaretçi döndürür, T tipinin sıfır değerinin değeri olan bellek adresini de döndürür.
**new her zaman işaretçi döndürür.**
make T(args) fonksiyonu new(T) fonksiyonundan farklı amaçlara sahiptir. make, slice, map ve channel için kullanılabilir ve başlangıç değeri olan bir T türü döndürür.
Bunun nedeni bu üç türün temel verilerinin bunlara işaret etmeden önce başlatılması gerektiğidir. Örneğin slice, temel diziye(underlying array), boyuta ve kapasiteye işaret eden bir işaretçi içerir. Bu veriler başlatılmadan önce, slice 'nil' dir, bu nedenle slice, map ve channel için make; temel verilerini başlatır ve bazı uygun değerleri atar.
**make sıfırdan farklı değerler döndürür**
Sıfır değeri boş anlamına gelmez. Çoğu zaman değişkenlerin varsayılan olarak aldığı değerdir.

```golang
int 0
int8 0
int32 0
int64 0
uint 0x0
rune 0 // rune'un asıl türü int32'dir
byte 0x0 // byte'ın asıl türü uint8'dir
float32 0 // boyut 4 byte'dır
float64 0 // boyut 8 byte'dır
bool false
string ""

```

__
