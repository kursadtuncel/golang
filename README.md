
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

Temel şekilde 'var' ile değişken atıyoruz. Go'da değişken türünü değişken adından sonra koyduğuna dikkat edin.

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
````

