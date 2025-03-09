package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go getProductDetailsFromExternalService(10)
	select {
	case productDetail := <-productChannel:
		fmt.Println("Ürün detayları getirildi", productDetail)
	case <-ctx.Done():
		fmt.Println("timeout occured, context cancelled")
	}
	fmt.Println("end of the main")
}

func getProductDetailsFromExternalService(productId int64) {
	time.Sleep(2 * time.Second)
	productChannel <- Product{
		Id:   productId,
		Name: "Televizyon",
	}
}
