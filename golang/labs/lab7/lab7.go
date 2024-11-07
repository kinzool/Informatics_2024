package lab7

import (
	"fmt"
	"log"
)

const ERROR = "Ошибка: "

type Product interface {
	applyDiscount(discount float64) error
	setPrice(newPrice float64) error
	getPrice() float64
	changeCharacteristics(name string, color string) error
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

func (e *Electronics) setPrice(newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	} else if newPrice == e.price {
		return fmt.Errorf("заданное значение равно исходному")
	} else {
		e.price = newPrice
		return nil
	}
}

func (c *Clothes) setPrice(newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	} else if newPrice == c.price {
		return fmt.Errorf("заданное значение равно исходному")
	} else {
		c.price = newPrice
		return nil
	}
}

func (f *Furniture) setPrice(newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	} else if newPrice == f.price {
		return fmt.Errorf("заданное значение равно исходному")
	} else {
		f.price = newPrice
		return nil
	}
}

func NewElectronic(name string, price float64, model string, color string) *Electronics {
	if price <= 0 {
		price = 1
	}
	e := &Electronics{name: name, price: price, model: model, color: color}
	return e
}

func NewClothes(name string, price float64, color string, size string) *Clothes {
	if price <= 0 {
		price = 1
	}
	c := &Clothes{name: name, price: price, color: color, size: size}
	return c
}

func NewFurniture(name string, price float64, color string) *Furniture {
	if price <= 0 {
		price = 1
	}
	f := &Furniture{name: name, price: price, color: color}
	return f
}

func (e *Electronics) applyDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	} else {
		e.price = e.price * (1 - (discount / 100))
		return nil
	}
}

func (c *Clothes) applyDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	} else {
		c.price = c.price * (1 - (discount / 100))
		return nil
	}
}

func (f *Furniture) applyDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	} else {
		f.price = f.price * (1 - (discount / 100))
		return nil
	}
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

func (e *Electronics) changeCharacteristics(name string, color string) error {
	if name == e.name || color == e.color {
		return fmt.Errorf("характеристики не поменялись")
	} else {
		e.name = name
		e.color = color
		return nil
	}
}
func (c *Clothes) changeCharacteristics(name string, color string) error {
	if name == c.name || color == c.color {
		return fmt.Errorf("характеристики не поменялись")
	} else {
		c.name = name
		c.color = color
		return nil
	}
}

func (f *Furniture) changeCharacteristics(name string, color string) error {
	if name == f.name || color == f.color {
		return fmt.Errorf("характеристики не поменялись")
	} else {
		f.name = name
		f.color = color
		return nil
	}
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
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------")
	product1 := NewElectronic("Macbook", 52000, "air m1 256gb", "gray")
	product2 := NewClothes("Свитшот", 2000, "Черный", "S")
	product3 := NewFurniture("Диван", 30000, "Фиолетовый")
	productsBeforeDiscount := []Product{product1, product2, product3}
	fmt.Printf("Цена до применения скидок: %.2f рублей\n", getSummaryAmount(productsBeforeDiscount))
	err := product1.applyDiscount(20)
	if err != nil {
		log.Fatal(fmt.Sprint(ERROR, err.Error()))
	}
	fmt.Printf("Цена первого товара после применения скидки: %.2f рублей\n", product1.getPrice())
	err = product2.setPrice(2500)
	if err != nil {
		log.Fatal(fmt.Sprint(ERROR, err.Error()))
	}
	err = product3.changeCharacteristics("Кровать", "Синий")
	if err != nil {
		log.Fatal(fmt.Sprint(ERROR, err.Error()))
	}
	err = product3.setPrice(27500)
	if err != nil {
		log.Fatal(fmt.Sprint(ERROR, err.Error()))
	}
	product3.printInformation()
	productsAfterDiscount := []Product{product1, product2, product3}
	fmt.Printf("Цена после применения скидки и изменения цены на третий товар: %.2f рублей\n", getSummaryAmount(productsAfterDiscount))
}
