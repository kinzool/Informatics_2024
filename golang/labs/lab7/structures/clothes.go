package structures

import "fmt"

type Clothes struct {
	name  string
	price float64
	color string
	size  string
}

func (c *Clothes) SetPrice(newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("заданная цена меньше нуля")
	} else if newPrice == c.price {
		return fmt.Errorf("заданное значение равно исходному")
	} else {
		c.price = newPrice
		return nil
	}
}

func NewClothes(name string, price float64, color string, size string) *Clothes {
	if price <= 0 {
		price = 1
	}
	c := &Clothes{name: name, price: price, color: color, size: size}
	return c
}

func (c *Clothes) ApplyDiscount(discount float64) error {
	if discount < 0 {
		return fmt.Errorf("заданная скидка меньше нуля")
	} else {
		c.price = c.price * (1 - (discount / 100))
		return nil
	}
}

func (c *Clothes) GetPrice() float64 {
	return c.price
}

func (c *Clothes) ChangeCharacteristics(name string, color string) error {
	if name == c.name || color == c.color {
		return fmt.Errorf("характеристики не поменялись")
	} else {
		c.name = name
		c.color = color
		return nil
	}
}
func (c *Clothes) PrintInformation() {
	fmt.Printf("name: %s, price: %.2f рублей, color: %s\n, size: %s", c.name, c.price, c.color, c.size)
}
