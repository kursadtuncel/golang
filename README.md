
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
[n]type'da n array'ın boyutudur, type ise array türüdür. Diğer dillerde olduğu gibi array değerlerini almak için [] kullanırız.

```golang
var arr [10]int
arr[0] = 42 // array 0 tabanlıdır
arr[1] = 13 // öğeye değer atama
fmt.Printf("The first element is %d\n", arr[0])
// öğe değerini çağırdık, 42 sayısını döndürecektir
fmt.Printf("The last element is %d\n", arr[9])
// 10. öğenin varsayılan değerini döndürecektir, ki bu sıfırdır.
```

Çünkü boyut(length) dizi türünün bir parçası olduğu için [3]int ve [4]int farklı türlerdir, bu nedenle dizilerin boyutunu değiştiremeyiz. Diziler argüman olarak kullanıldığında, fonksiyonlar referanslar yerine kopyalarını alır. Referansları kullanmak istiyorsak Slice kullanmak isteyebilirsiniz.

## slice
Birçok senaryoda diziler iyi bir tercih olmayabilir. Özellikle diziyi tanımlarken boyutunun ne olacağını bilmediğimizde. Bu nedenle daha dinamik bir diziye ihtiyacımız vardır. Buna slice denir.
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
- 'len' map için de çalışır, map'in kaç tane s anahtarına sahip olduğunu döndürür.
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

# Control Statements and Functions
## if

```golang
// go'da if kullanırken paranteze gerek yoktur.
if x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is less than equal to 10")
}
```
Go'da if ile ilgili en kullanışlı şey, if'den önce başlatma ifadesi tanımlayabilmemizdir.
(Yalnızca if bloğunun içinde kullanılabilir.)

```golang
if x := computedValue(); x > 10{
    fmt.Println("x is the greater than 10")
} else {
    fmt.Println("x is less than 10")
}

fmt.Println(x)
```
```golang
if integer == 3 {
    fmt.Println("the integer is equal to 3")
} else if integer < 3 {
    fmt.Println("the integer is less than 3")
} else {
    fmt.Println("the integer is greater than 3")
}
```


## goto
goto, kontrol akışını daha önce tanımlanmış bir etikete yönlendirir. Ancak aynı kod bloğu içerisinde kullanılırken dikkatli olmalısınız.

```golang
func myFunc() {
    i := 0
    Here: // etiket : ile bitiyor
    fmt.Println(i)
    i++
    goto Here 
    // etiket adı büyük-küçük harflere duyarlıdır
}
```
## for
```golang
for expresssion1; expression2; expression3 {
    //...
}
```

Burada expression1 ve expression3 değişken tanımları ve fonksiyondan dönen değerlerdir ve expression2 bir koşullu ifadedir. expression1 döngüden önce bir kez yürütülecek, expression3 ise her döngüden sonra yürütülecektir.

```golang
package main

import "fmt"

func main(){
    sum := 0;
    for index:=0; index < 10 ; index++{
        sum += index
    }
    fmt.Println("sum is equal to ", sum)
    // sum is equal to 45
}
```

- Birden fazla atamaya ihtiyaç duyarsak, Go'da bunun operatörü yoktur. Bu nedenle i, j =
i + 1, j - 1 gibi paralel atamalar kullanırız.
gerekli değilse expression1 ve expression3'ü atlayabiliriz

```golang
sum := 1
for ; sum < 1000; {
    sum += sum
}
```
;'yi nasıl atladığımızı gördünüz mü? while ile aynı!

```golang
sum := 1
for sum < 1000 {
    sum += sum
}
```
break, continue operatörleri:

```golang
for index := 10; index>0; index-- {
    if index == 5{
        break // veya continue
    }
    fmt.Println(index)
    // break prints 10、9、8、7、6
    // continue prints 10、9、8、7、6、4、3、2、1
}
```

for, range kullanıldığında array, slice, map'ten veri okuyabilir.
```golang
for k,v := range map{
    fmt.Println("map's key: ", k)
    fmt.Printlln("map's val: ", v)
}
```

## switch
Çok fazla if-else kullanma sorununu çözebilmek adına switch kullanmak yararlı olabilir.
```golang
switch sExpr {
    case expr1:
    some instructions
    case expr2:
    some other instructions
    case expr3:
    same other instructions
    default: 
    other code
}
```
Koşullar sabit olmak zorunda değildir ve eşleşene kadar işlem yukarıdan aşağıya doğru yürütülür. Ayrıca switch anahtarından sonra bir ifade yoksa, o zaman **true** ile eşleşir.

```golang
i := 10
switch i {
    case 1:
    fmt.Println("i is equal to 1")
    case 2,3,4:
    fmt.Println("i is equal to 2,3,4")
    case 10:
    fmt.Println("i is equal to 10")
    default:
    fmt.Println("all i know is that i is an integer")
    // 5. satırda (case 2,3,4:) birçok değer koyduk ve case'in gövdesinin sonuna break eklememize gerek yok. herhangi bir case ile eşleştiğinde switch gövdesinden dışarı atlayacaktır.
}
```

Daha fazla case ile eşleşmeye devam etmek istiyorsak, **fallthrough** ifadesini kullanmamız gerekir.

```golang
integer := 6
switch integer {
    case 4:
        fmt.Println("integer <=4")
        fallthrough
        fmt.Println("integer <= 5")
        fallthrough
    case 6:
        fmt.Println("integer <= 6")
        fallthrough
    case 7:
        fmt.Println("integer <= 7")
        fallthrough
    case 8:
        fmt.Println("integer <= 8")
        fallthrough
    default:
        fmt.Println("default case")
}
```

Çıktı şöyle olacaktır:
```
integer <= 6
integer <= 7
integer <= 8
default case
```
# Functions (Fonksiyonlar)

```golang
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
    //function body
    //multi-value return
    return value1, value2
}
```

Yukarıdaki örnekten şu bilgileri çıkarmalıyız:
- Bir fonksiyon adı tanımlamak için **func** kullanıyoruz.
- Fonksiyonların sıfır, bir veya birden fazla argümanı vardır. Argüman türü argüman adından sonra gelir ve argümanlar "," ile ayrılır.
- Fonksiyonlar birden fazla değer döndürebilir.
- output1 ve output2 adında iki adet dönüş değeri var, bunların isimlerini yazmayıp sadece türlerini kullanabiliriz.
- Eğer sadece bir tane dönüş değeri varsa ve ismi atladıysak, dönüş değerleri için parantez kullanmamıza gerek yoktur.
- Fonksiyonun dönüş değerleri yoksa, dönüş parametrelerini tamamen atlayabiliriz.
- Fonksiyonun dönüş değerleri varsa, return ifadesini fonksiyonun gövdesinde kullanmalıyız.

```golang
package main

import "fmt"

// A+B ve A*B nin sonuçlarını döndüreceğiz:

func SumAndProduct(A, B int) (int, int) {
    return A + B,  A* B
}

func main(){
    x := 3
    y := 4

    xPLUSy, xTIMESy := SumAndProduct(x, y)
    fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
    fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
```
Eğer fonksiyonları paketin dışında kullanacaksak, fonksiyon isimleri büyük harfle başlıyorsa, return için tam ifadeler yazmak daha iyi olur, bu kodları daha okunabilir hale getirir.

```golang
func SumAndProduct(A,B,int) (add int, multiplied int) {
    add = A+B
    multiplied = A*B
    return
}
```

## Variadic Functions (Değişken Fonksiyonlar)
Go, değişken sayıda argümana sahip fonksiyonları destekler. Bu fonksiyonlara "variadic" denir, bu da fonksiyonun belirsiz sayıda argümana izin verdiği anlamına gelir. 

````
func myFunc(arg ...int) {}
````

arg ...int, Go'ya bu değişkenlerin argümanları olan bir fonksiyon olduğunu söyler. 

```golang
for _, n := range arg {
    fmt.Printf("And the number is : %d\n", n)
}
```

## Pass by value and pointers

Çağırılan fonksiyona bir argüman verdiğimizde, o fonksiyon aslında değişkenlerimizin kopyasını alır, dolayısıyla herhangi bir değişiklik orijinal değişkeni etkilemez.

```golang
package main

import "fmt"

// a'ya 1 eklediğimiz basit bir fonksiyon:
func add1(a int) int{
    a = a+1 //a nın değerini değiştiriyoruz
    return a //a nın yeni değerini döndürüyoruz
} 

func main(){
    x := 3
    fmt.Println("x = ", x) // x=3 çıktısı verir

    x1 := add1(x) //add1(x)'i çağırır

    fmt.Println("x+1 =", x1) //x+1=4 çıktısı verir
    fmt.Println("x = ", x) //x = 3 çıktısı verir
}
```

Görüldüğü üzere, add1'i x ile çağırdığımız halde x'in başlangıç değeri değişmiyor. Sebebi ise çok basit: add1'i çağırdığımızda ona x'in kendisini değil, bir kopyasını verdik.

Şimdi gerçek x'i fonksiyona nasıl geçirebileceğeimizi anlayalım.
Burada işaretçileri kullanmamız gerekiyor. Değişkenlerin bellekte sakladığı ve bazı bellek adreslerine sahip olduğunu biliyoruz. Yani, bir değişkenin değerini değiştirmek istiyorsak, bellek adresini değiştirmeliyiz. Bu nedenle add1 fonksiyonu değerini değiştirmek için x'in bellek adresini bilmelidir. Burada fonksiyona %x iletiyoruz ve argümanın türünü *int işaretçi türü olarak değiştiriyoruz. Değerin bir kopyasını değil, işaretçinin bir kopyasını ilettiğimizi unutmayın.

```golang
package main

import "fmt"

func add1(a *int) int {
    *a = *a + 1
    return a
}

func main() {
    x := 3
    fmt.Println("x = ", x)
    x1 := add1(%x)
    fmt.Println("x+1 =", x1) // x+1 = 4 çıktısını verir
    fmt.Println("x =", x) // x = 4 çıktısını verir.
}
```
Şimdi x'in değerini fonksiyonlarda değiştirebiliyoruz.

## defer
Go'da defer adında çok iyi tasarlanmış bir anahtar sözcük vardır. Bir fonksiyonda birçok defer ifadesi olabilir, program fonksiyonların sonuna kadar yürütüldüğünde bunlar ters sırada yürütülür. Programın bazı kaynak dosyalarını açtığı durumda, fonksiyon hatalarla geri dönebilmeden önce bu dosyaların kapatılması gerekir. Örnek üzerinden inceleyelim:

```golang
func ReadWrite() bool {
    file.Open("file")
    // some code
    if failureX {
        file.Close()
        return false
    }

    if failureY {
        file.Close()
        return false
    }

    file.Close()
    return true
}
```
Görüldüğü üzere kodları tekrarlamak zorundayız. defer kullanarak daha temiz ve okunabilir kod yazabiliriz:

```golang
func ReadWrite() bool{
    file.Open("file")
    defer file.Close()
    if failureX {
        return false
    }
    if failureY {
        return false
    }
    return true
}
```
Birden fazla defer varsa, kodlar ters sırayla yürütülür. Aşağıdaki örnekte 4 3 2 1 0 çıktısı alınır:

```golang
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d", i)
}
```
## Functions as values and types
Go'da fonksiyonlar da birer değişkendir ve bunları tanımlamak için **type** kullanılır.
Aynı imzaya sahip fonksiyonlar aynı tür olarak görülebilir. 

```golang
type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
```
Bu özelliğin avantajı nedir diye soracak olursanız, bunun cevabı fonksiyonları değer olarak geçirmemize izin vermesidir.

```golang
package main

import "fmt"

type testInt func(int) bool // değişken tipinde bir fonksiyon tanımladık

func isOdd(integer int) bool {
    return integer%2 != 0
}

func isEven(integer int) bool {
    return integer%2 == 0
}

// 'f' fonksiyonunu başka bir fonksiyona argüman olarak geçiriyoruz
func filter(slice []int, f testInt) []int {
    var result []int
    for _, value := range slice {
        if f(value) {
            result = append(result, value)
        }
    }
    return result
}

var slice = []int{1,2,3,4,5,7}

func main(){
    odd := filter(slice, isOdd)
    even := filter(slicee, isEven)

    fmt.Println("slice =", slice)
    fmt.Println("Odd elements of slice are", odd),
    fmt.Println("Even elements of slice are", even)
}

```

Arayüzleri kullandığımızda çok işe yarayacaktır. Örnekte görüldüğü üzere testInt, tür olarak bir fonksiyona sahip bir değişkendir ve filter'in döndürülen değerleri ve argümanları testInt ile aynıdır. Dolayısıyla programlarımızda karmaşık mantığa sahip olabilirken, kodda esnekliği koruyabiliriz.

## Panic and Recover
Go, Java gibi try-catch yapısına sahip değildir. Go hatalarla başa çıkabilmek için istisnalar yerine panic ve recover kullanır. Ancak panic çok güçlü olmasına rağmen çok fazla kullanmamaya dikkat etmeliyiz.
Panic programların normal akışını bozup panik durumuna geçmeyi sağlayan yerleşik bir fonksiyondur. "F" fonksiyonu panic'i çağırdığında, "F" çalışmaya devam etmez ancak defer fonksiyonları çalışmaya devam eder. Ardından "F" panik durumunun oluştuğu noktaya geri döner.
Program tüm bu fonksiyonlar panikle **goroutine**'in ilk seviyesine dönene kadar sonlanmayacaktır. Panic, program içerisinde "panic" çağırılarak üretilebildiği gibi, array erişiminin sınır dışına çıkması gibi bazı hatalar da panic oluşumuna sebep olur.

Recover, goroutine s'i panic durumundan kurtarmak için kullanılan yerleşik bir fonksiyondur. defer fonksiyonlarında recover'ı çağırmak faydalıdır çünkü program panik durumundayken normal fonksiyonlar yürütülmeyecektir. Program panic durumunda ise panic değerlerini yakalar, panic durumunda değilse "nil" değerini alır.

```golang
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```
panic durumunu aşağıdaki örnekte olduğu gibi kontrol edebiliriz.
```golang
func throwsPanic(f func()) (b bool) {
    defer func() {
        if x := recover(); x != nil {
            b = true
        }()
        f() // eğer if panic'e neden olursa, recover devreye girecektir.
        return
    }
}
```

## main and init functions
Go'nun main ve init adı verilen iki tutma biçmi vardır; init tüm paketlerde kullanılabilirken, main yalnızca main paketinde kullanılabilir. Bu iki fonksiyonun argümanları olamaz ve değer döndüremezler.
- Bir pakette birden fazla init fonksiyonu yazabiliriz, ancak her paket için yalnızca bir adet init fonksiyonu yazılması daha doğrudur.
Go programları main() ve init()'i otomatik olarak çağırır, bu yüzden bizim çağırmamıza gerek kalmaz. Her paket için init fonksiyonu isteğe bağlıdır, ancak main paketinin yalnızca bir tane main fonksiyonu vardır. Programlar main paketinden başlatılır ve yürütülür. main paketi diğer paketleri içe aktarırsa, derleme sırasında içe aktarma gerçekleşir. Bir paket birçok kez içe aktarılırsa, yalnızca bir kez derlenir. Paketler içe aktarıldıktan sonra, programlar içe aktarılan paketlerdeki sabitleri ve değişkenleri başlatır, ardından varsa init fonksiyonu yürütülür ve program bu şekilde devam eder. 

## import
```golang
import (
    "fmt"
)

fmt.Printline("hello world")
```
fmt, Go standart kütüphanesindendir, $GOROOT/pkg içerisinde bulunur. Go, üçüncü taraf paketlerini iki şekilde destekler:

- 1: Göreceli olarak içe aktarma "./model" // paketi aynı dizine yükler, önerilmez
- 2: Mutlak olarak içe aktarma "shorturl/model" paketi "$GOPATH/pkg/shorturl/model" //önerilir

# struct
Go'da daha karmaşık bir veri yapısını tek değişkende barındırabilmek için struct kullanılır.
```golang
type person struct{
    name string
    age int
}
```

```golang
type person struct {
    name string
    age int
}

var P person // p person'un türüdür.

P.name = "Sarah"
P.age = 25
fmt.Printf("The person's name is %s\n", P.name)
```

Struct tanımlamanın 3 ayrı yolu daha vardır:

```
- P := person{"Tom", 25}
- P := person{age:24, name:"Bob"}
- P := struct{name string; age int}{"Amy", 18}
```

```golang
package main

import "fmt"

type person struct {
    name string
    age int
}

func Older(p1, p2 person) (person, int) {
    if p1.age > p2.age {
        return p1, p1.age - p2.age
    }
    return p2, p2.age - p1.age
}

func main() {
    var tom person
    tom.name, tom.age = "Tom", 18
    bob := person{age:25, name: "Bob"}
    paul := person{"Paul", 43}

    tb_Older, tb_diff := Older(tom, bob)
    tb_Older, tb_diff := Older(tom, paul)
    bp_Older, bp_diff := Older(bob, paul)
    
    fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)
}
```