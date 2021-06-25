package component

import "fmt"

type Bevarage interface {
	//计算价格
	Cost() int
	//返回描述
	Me() string
}
type Tea struct {
	Bevarage
	name string
	price int
}


func (self *Tea) Me() string {
	return self.name
}
func (self *Tea) Cost() int {
	return self.price
}

type Moli struct {
	*Tea
}

func NewMoli() *Moli {
	return &Moli{&Tea{name: "茉莉",price: 18}}
}

func (self *Moli) Me() string {
	return self.name
}
func (self *Moli) Cost() int{
	return self.price
}
type Puer struct {
	*Tea
}

func NewPuer() *Puer {
	return &Puer{&Tea{name: "普洱", price: 38}}
}

func (self *Puer) Me() string{
	return self.name
}

func (self *Puer) Cost() int {
	return self.price
}

func main() {
	//结构体类型相同，在内存中是连续存储的

	moli:=NewMoli()
	puer:=NewPuer()
	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n", 1, moli.Me(), moli.Cost())
	fmt.Printf("第 %v 杯是 %s 售价 %v 元\n", 2, puer.Me(),puer.Cost())

	
}
