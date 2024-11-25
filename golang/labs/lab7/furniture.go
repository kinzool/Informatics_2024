package lab7

import (
	"fmt"
)

type Furniture struct {
	name  string
	price float64
	color string
}

func NewFurniture(name string, price float64, color string) *Furniture {
	if price <= 0 {
		price = 1
	}
	f := &Furniture{name: name, price: price, color: color}
	return f
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

func (f *Furniture) applyDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	} else if discount > 100 {
		return fmt.Errorf("заданная скидка больше ста")
	}
	f.price = f.price * (1 - (discount / 100))
	return nil
}

func (f *Furniture) getPrice() float64 {
	return f.price
}

func (f *Furniture) printInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, color: %s\n", f.name, f.price, f.color)
}

func (f *Furniture) ChangeFurnitureColor(newColor string) error {
	if newColor == f.color {
		return fmt.Errorf("цвет не изменился")
	}
	f.color = newColor
	return nil
}
