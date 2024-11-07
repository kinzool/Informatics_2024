package lab7

import (
	"fmt"
)

type Product interface {
	applyDiscount(discount float64)
	setPrice(newPrice float64)
	getPrice() float64
	changeCharacteristics(name string, color string)
	printInformation()
}

type Electronics struct {
	name  string
	price float64
	model string
	color string
}

type Clothes struct {
	name  string
	price float64
	color string
	size  string
}

type Furniture struct {
	name  string
	price float64
	color string
}

func (e *Electronics) setPrice(newPrice float64) {
	e.price = newPrice
}

func (c *Clothes) setPrice(newPrice float64) {
	c.price = newPrice
}

func (f *Furniture) setPrice(newPrice float64) {
	f.price = newPrice
}

func NewElectronic(name string, price float64, model string, color string) *Electronics {
	e := &Electronics{name: name, price: price, model: model, color: color}
	return e
}

func NewClothes(name string, price float64, color string, size string) *Clothes {
	c := &Clothes{name: name, price: price, color: color, size: size}
	return c
}

func NewFurniture(name string, price float64, color string) *Furniture {
	f := &Furniture{name: name, price: price, color: color}
	return f
}

func (e *Electronics) applyDiscount(discount float64) {
	e.price = e.price * (1 - (discount / 100))
}

func (c *Clothes) applyDiscount(discount float64) {
	c.price = c.price * (1 - (discount / 100))
}

func (f *Furniture) applyDiscount(discount float64) {
	f.price = f.price * (1 - (discount / 100))
}

func (e *Electronics) getPrice() float64 {
	return e.price
}

func (c *Clothes) getPrice() float64 {
	return c.price
}

func (f *Furniture) getPrice() float64 {
	return f.price
}

func (e *Electronics) changeCharacteristics(name string, color string) {
	e.name = name
	e.color = color
}
func (c *Clothes) changeCharacteristics(name string, color string) {
	c.name = name
	c.color = color
}

func (f *Furniture) changeCharacteristics(name string, color string) {
	f.name = name
	f.color = color
}

func (e *Electronics) printInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, model: %s\n, color: %s", e.name, e.price, e.model, e.color)
}

func (c *Clothes) printInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, color: %s\n, size: %s", c.name, c.price, c.color, c.size)
}

func (f *Furniture) printInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, color: %s\n", f.name, f.price, f.color)
}

func getSummaryAmount(products []Product) float64 {
	var summaryAmount float64 = 0
	for _, product := range products {
		summaryAmount += product.getPrice()
	}
	return summaryAmount
}

func RunLab7() {
	product1 := NewElectronic("Macbook", 52000, "air m1 256gb", "gray")
	product2 := NewClothes("Свитшот", 2000, "Черный", "S")
	product3 := NewFurniture("Диван", 30000, "Фиолетовый")
	productsBeforeDiscount := []Product{product1, product2, product3}
	fmt.Printf("Цена до применения скидок: %.2f рублей\n", getSummaryAmount(productsBeforeDiscount))
	product1.applyDiscount(20)
	fmt.Printf("Цена первого товара после применения скидки: %.2f рублей\n", product1.getPrice())
	product2.setPrice(2500)
	product3.changeCharacteristics("Кровать", "Синий")
	product3.setPrice(27500)
	product3.printInformation()
	productsAfterDiscount := []Product{product1, product2, product3}
	fmt.Printf("Цена после применения скидки и изменения цены на третий товар: %.2f рублей\n", getSummaryAmount(productsAfterDiscount))
}
