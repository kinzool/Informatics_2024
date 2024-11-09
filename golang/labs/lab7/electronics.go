package lab7

import "fmt"

type Electronics struct {
	name  string
	price float64
	model string
	color string
}

func (e *Electronics) setPrice(newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	} else if newPrice == e.price {
		return fmt.Errorf("заданное значение равно исходному")
	}
	e.price = newPrice
	return nil
}

func NewElectronic(name string, price float64, model string, color string) *Electronics {
	if price <= 0 {
		price = 1
	}
	e := &Electronics{name: name, price: price, model: model, color: color}
	return e
}

func (e *Electronics) applyDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	}
	e.price = e.price * (1 - (discount / 100))
	return nil
}

func (e *Electronics) getPrice() float64 {
	return e.price
}

func (e *Electronics) changeCharacteristics(name string, color string) error {
	if name == e.name || color == e.color {
		return fmt.Errorf("характеристики не поменялись")
	}
	e.name = name
	e.color = color
	return nil
}

func (e *Electronics) printInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, model: %s\n, color: %s", e.name, e.price, e.model, e.color)
}