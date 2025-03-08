package main

import (
	"fmt"
)

func main() {
	productService := ProductService{}
	err := productService.Add(Product{id: 1, name: "", price: 3000})
	if err != nil {
		fmt.Println(err)
	}
}

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct{}

func (ProductService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{text: "Urun ismi bos olamaz", code: "1001"}
	}
	if product.price < 0 {
		return ValidationError{text: "Urun fiyatı 0'dan buyuk olmalıdır", code: "2001"}
	}
	fmt.Println("ürün eklendi!")
	return nil
}

type ValidationError struct {
	code string
	text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Hata : %s, Kod: %s", validationError.text, validationError.code)
}
