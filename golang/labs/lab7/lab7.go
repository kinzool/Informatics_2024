package lab7

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/labs/lab7/structures"
)

const ERROR = "Ошибка: "

type Product interface {
	ApplyDiscount(discount float64) error
	SetPrice(newPrice float64) error
	GetPrice() float64
	ChangeCharacteristics(name string, color string) error
	PrintInformation()
}

func getSummaryAmount(products []Product) float64 {
	var summaryAmount float64 = 0
	for _, product := range products {
		summaryAmount += product.GetPrice()
	}
	return summaryAmount
}

func RunLab7() {
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------")
	product1 := structures.NewElectronic("Macbook", 52000, "air m1 256gb", "gray")
	product2 := structures.NewClothes("Свитшот", 2000, "Черный", "S")
	product3 := structures.NewFurniture("Диван", 30000, "Фиолетовый")
	productsBeforeDiscount := []Product{product1, product2, product3}
	fmt.Printf("Цена до применения скидок: %.2f рублей\n", getSummaryAmount(productsBeforeDiscount))
	err := product1.ApplyDiscount(20)
	if err != nil {
		log.Fatal(fmt.Sprint(ERROR, err.Error()))
	}
	fmt.Printf("Цена первого товара после применения скидки: %.2f рублей\n", product1.GetPrice())
	err = product2.SetPrice(2500)
	if err != nil {
		log.Fatal(fmt.Sprint(ERROR, err.Error()))
	}
	err = product3.ChangeCharacteristics("Кровать", "Синий")
	if err != nil {
		log.Fatal(fmt.Sprint(ERROR, err.Error()))
	}
	err = product3.SetPrice(27500)
	if err != nil {
		log.Fatal(fmt.Sprint(ERROR, err.Error()))
	}
	product3.PrintInformation()
	productsAfterDiscount := []Product{product1, product2, product3}
	fmt.Printf("Цена после применения скидки и изменения цены на третий товар: %.2f рублей\n", getSummaryAmount(productsAfterDiscount))
}
