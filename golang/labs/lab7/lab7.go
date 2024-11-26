package lab7

import (
	"fmt"
	"log"
)

const ERRORPASTE = "Ошибка: "

type Product interface {
	applyDiscount(discount float64) error
	getPrice() float64
	printInformation()
}

func getSummaryAmount(products []Product) float64 {
	var summaryAmount float64 = 0
	for _, product := range products {
		summaryAmount += product.getPrice()
	}
	return summaryAmount
}

func setPrice(product interface{}, newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	}
	switch p := product.(type) {
	case *Electronics:
		if newPrice == p.price {
			return fmt.Errorf("заданное значение равно исходному")
		}
		p.price = newPrice
	case *Clothes:
		if newPrice == p.price {
			return fmt.Errorf("заданное значение равно исходному")
		}
		p.price = newPrice
	case *Furniture:
		if newPrice == p.price {
			return fmt.Errorf("заданное значение равно исходному")
		}
		p.price = newPrice
	default:
		return fmt.Errorf("неизвестный тип продукта")
	}
	return nil
}

func RunLab7() {
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------")
	product1 := NewElectronic("Macbook", 52000, "air m1 256gb", "gray")
	product2 := NewClothes("Свитшот", 2000, "Черный", "S")
	product3 := NewFurniture("Диван", 30000, "Фиолетовый")
	productsBeforeDiscount := []Product{product1, product2, product3}
	fmt.Printf("Цена до применения скидок: %.2f рублей\n", getSummaryAmount(productsBeforeDiscount))
	err := product1.applyDiscount(20)
	if err != nil {
		log.Fatal(fmt.Sprint(ERRORPASTE, err.Error()))
	}
	fmt.Printf("Цена первого товара после применения скидки: %.2f рублей\n", product1.getPrice())
	err = setPrice(product2, 2500)
	if err != nil {
		log.Fatal(fmt.Sprint(ERRORPASTE, err.Error()))
	}
	err = product3.ChangeFurnitureColor("Красный")
	if err != nil {
		log.Fatal(fmt.Sprint(ERRORPASTE, err.Error()))
	}
	err = product1.ChangeElectronicsModel("pro m2 512gb")
	if err != nil {
		log.Fatal(fmt.Sprint(ERRORPASTE, err.Error()))
	}
	err = product2.ChangeClothesSize("XL")
	if err != nil {
		log.Fatal(fmt.Sprint(ERRORPASTE, err.Error()))
	}
	err = setPrice(product3, 27500)
	if err != nil {
		log.Fatal(fmt.Sprint(ERRORPASTE, err.Error()))
	}
	product3.printInformation()
	productsAfterDiscount := []Product{product1, product2, product3}
	fmt.Printf("Цена после применения скидки и изменения цены на третий товар: %.2f рублей\n", getSummaryAmount(productsAfterDiscount))
}
