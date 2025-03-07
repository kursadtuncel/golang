package main

import "fmt"

type IShippable interface {
	shippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

type Flower struct {
	desi int
}

func (book *Book) shippingCost() int {
	return 5 + book.desi*2
}

func (electronic *Electronic) shippingCost() int {
	return 10 + electronic.desi*3
}

func (flower *Flower) shippingCost() int {
	return 12 + flower.desi*3
}

func main() {
	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Book{desi: 20},
		&Electronic{desi: 20},
		&Flower{desi: 10},
	}

	fmt.Println(calculateTotalShippingCost(products))
	/*var product IShippable
	product = &Book{desi: 10}

	fmt.Println(product.shippingCost())
	product = &Electronic{desi: 10}
	fmt.Println(product.shippingCost())
	/*
		book1 := &Book{desi: 10}
		book2 := &Book{desi: 20}

		fmt.Println(book1.shippingCost())
		fmt.Println(book2.shippingCost())
	*/
	/*
		electronics := []Electronic{{desi: 10}, {desi: 20}}
		calculateTotalShippingCostOfElectronics(electronics)
		books := []Book{{desi: 10}, {desi: 20}}
		fmt.Println(calculateTotalShippingCostOfBooks(books))
	*/
}

func calculateTotalShippingCost(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.shippingCost()
	}
	return total
}

/*
func calculateTotalShippingCostOfBooks(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.shippingCost()
	}
	return total

}

func calculateTotalShippingCostOfElectronics(electronics []Electronic) int {
	total := 0
	for _, electronic := range electronics {
		total += electronic.shippingCost()
	}
	return total
}
*/
