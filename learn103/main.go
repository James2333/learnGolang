package main

import (
	"fmt"
)

//抽象构建接口
//定义了抽象构件接口“饮料”接口，它包含了两个方法，输出价格和描述自己，饮料接口作为最底层接口是所有饮料都必须要实现的
type Beverage interface {
	//计算价格
	Cost() int
	//返回描述
	Me() string
}

//定义一个具体的产品Tea，包含Beverage
type Tea struct {
	Beverage
	Price int
	Name  string
}

func (self *Tea) Me() string {
	return self.Name
}
func (self *Tea) Cost() int {
	return self.Price
}

type Moli struct {
	*Tea
}

func NewMoli() *Moli {
	return &Moli{&Tea{Name: "茉莉", Price: 48}}
}

func (self *Moli) Me() string {
	return self.Name
}

func (self *Moli) Cost() int {
	return self.Price
}

type Puer struct {
	*Tea
}

func NewPuer() *Puer {
	return &Puer{&Tea{Name: "普洱", Price: 38}}
}

func (self *Puer) Me() string {
	return self.Name
}

func (self *Puer) Cost() int {
	return self.Price
}

type Condiment struct {
	*Tea     //作用？
	beverage Beverage
	name     string
	price    int
}

func (self *Condiment) Me() string {
	return self.name
}

func (self *Condiment) Cost() int {
	return self.price
}

type Sugar struct {
	*Condiment
}

func NewSugar(beverage Beverage) *Sugar {
	return &Sugar{&Condiment{beverage: beverage, name: "糖", price: 3}}
}

func (self *Sugar) Me() string {
	return self.beverage.Me() + " 加点 " + self.name
}

func (self *Sugar) Cost() int {
	return self.beverage.Cost() + self.price
}

type Ice struct {
	*Condiment
}

func NewIce(beverage Beverage) *Ice {
	return &Ice{&Condiment{beverage: beverage, name: "冰", price: 3}}
}

func (self *Ice) Me() string {
	return "加了" + self.name + "的" + self.beverage.Me()
}

func (self *Ice) Cost() int {
	return self.beverage.Cost() + self.price
}

func main() {
	moli := NewMoli()
	puer := NewPuer()

	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n", 1, moli.Me(), moli.Cost())
	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n", 2, puer.Me(), puer.Cost())

	fmt.Printf("下面我们给刚才那杯茉莉加点糖...\n")
	sugar := NewSugar(moli)
	fmt.Printf("刚刚给茉莉加了点糖，现在准备尝一下\n")
	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n", 3, sugar.Me(), sugar.Cost())

	ice := NewIce(puer)
	fmt.Printf("来一杯加冰的普洱，现在准备尝一下\n")
	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n", 4, ice.Me(), ice.Cost())
	fmt.Printf("好喝吗，欢迎再来 ^_^ ")
}
