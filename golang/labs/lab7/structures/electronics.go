package structures

import "fmt"

type Electronics struct {
	name  string
	price float64
	model string
	color string
}

func (e *Electronics) SetPrice(newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	} else if newPrice == e.price {
		return fmt.Errorf("заданное значение равно исходному")
	} else {
		e.price = newPrice
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

func (e *Electronics) ApplyDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	} else {
		e.price = e.price * (1 - (discount / 100))
		return nil
	}
}

func (e *Electronics) GetPrice() float64 {
	return e.price
}

func (e *Electronics) ChangeCharacteristics(name string, color string) error {
	if name == e.name || color == e.color {
		return fmt.Errorf("характеристики не поменялись")
	} else {
		e.name = name
		e.color = color
		return nil
	}
}

func (e *Electronics) PrintInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, model: %s\n, color: %s", e.name, e.price, e.model, e.color)
}
