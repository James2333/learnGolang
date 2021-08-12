package main

import (
	"fmt"
	"learn101/mock_interface_test/equipment"
)

type Person struct {
	name  string
	phone equipment.Phone
}

func NewPerson(name string, phone equipment.Phone) *Person {
	return &Person{
		name:  name,
		phone: phone,
	}
}

func (x *Person) goSleep() {
	fmt.Printf("%s go to sleep!", x.name)
}

func (x *Person) dayLife() bool {
	fmt.Printf("%s's daily life:\n", x.name)
	if x.phone.WeiXin() && x.phone.WangZhe() && x.phone.ZhiHu() {
		x.goSleep()
		return true
	}
	return false
}