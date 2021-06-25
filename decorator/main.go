package decorator

import (
	"learn101/component"
)

type Condiment struct {
	*component.Tea
	beverage component.Bevarage
	name string
	price int
}

func (self *Condiment) Me() string{
	return self.name
}

func (self *Condiment) Cost() int {
	return self.price

}

type Sugar struct {
	*Condiment
}

func NewSugar(bevarage component.Bevarage) *Sugar {
	return &Sugar{&Condiment{beverage: bevarage,name: "糖",price: 3}}
	
}
func (self *Sugar) Me() string{
	return self.beverage.Me()+"加点"+self.name

}
func (self *Sugar) Cost() int {
	return self.price+self.beverage.Cost()

}

type Ice struct {
	*Condiment
}

func NewIce(beverage component.Bevarage) *Ice {
	return &Ice{ &Condiment{beverage: beverage, name: "冰", price: 3 }}
}

func (self *Ice) Me() string {
	return "加了" + self.name + "的" + self.beverage.Me()
}

func (self *Ice) Cost() int {
	return self.beverage.Cost() + self.price
}