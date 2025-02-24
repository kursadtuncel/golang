
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

Derleme sırasında belirlenen değerlerdir ve runtime sürecinde değiştiremeyeceğimiz değerlerdir.
Go'da sabit türleri olarak sayı, boolean veya dize kullanabiliriz.

````golang
const constantName = value
// gerektiğinde sabit türlerini atlayabiliriz
const Pi float32= 3.1415926
```
