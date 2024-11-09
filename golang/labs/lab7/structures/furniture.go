package structures

import "fmt"

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
func (f *Furniture) SetPrice(newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	} else if newPrice == f.price {
		return fmt.Errorf("заданное значение равно исходному")
	} else {
		f.price = newPrice
		return nil
	}
}

func (f *Furniture) ApplyDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	} else {
		f.price = f.price * (1 - (discount / 100))
		return nil
	}
}

func (f *Furniture) GetPrice() float64 {
	return f.price
}

func (f *Furniture) ChangeCharacteristics(name string, color string) error {
	if name == f.name || color == f.color {
		return fmt.Errorf("характеристики не поменялись")
	} else {
		f.name = name
		f.color = color
		return nil
	}
}

func (f *Furniture) PrintInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, color: %s\n", f.name, f.price, f.color)
}
