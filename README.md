
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
Go'da aynı anda hem **int** hem de **uint** vardır, aynı uzunluğa sahiptirler, ancak
belirli uzunluk işletim sisteminize bağlıdır.
32-bit işletim sistemlerinde 32bit, 64bit işletim sistemlerinde 64bit kullanılır.
Go'da ayrıca rune, int8, int16, int32, int64, byte, uint8, uint16, uint32, uint64 gibi belirli uzunluklara sahip tipler de vardır. rune'un int32'nin, byte'ın ise uint8'in takma adı(alias) olduğunu unutmayın.

Bilmeniz gereken en önemli şey, bu değerler arasında değer atayamayacağınızdır, böyle bir işlem derleme hatasına yol açacaktır.

````
var a int8
var b int32
c := a + b
````
int32, int8'den daha uzun bir uzunluğa sahip olmasına ve int ile aynı türe sahip olmasına rağmen, aralarında değer atayamazsınız. (c burada int türü olarak doğrulanacaktır)

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
````golang
var s string = "Hello"
s[0] = 'c'
```