package main

import "learn101/mock_interface_test/equipment"

func main() {
	phone:=equipment.NewIphone6s()
	xm:=NewPerson("xiaoming",phone)
	xm.dayLife()
}
