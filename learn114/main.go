package main

import (
	"fmt"
	"sync"
)

type Car interface {
	Run()
}

type Dirver interface {
	Dirve(car Car)
}

type Benz struct {

}

func (benz *Benz)Run()  {
	fmt.Println("benz si running")
}

type Zhang3 struct {

}

func (zhang3 *Zhang3)Drive(car Car)  {
	fmt.Println("zhang3 dirve car")
	car.Run()
}

func main() {
	//benz:=&Benz{}
	//
	//zhang3:=&Zhang3{}
	//
	//zhang3.Drive(benz)
	var m sync.Map
	//m.Load("zbc")
	m.Store("zbc",make(map[string]int))
	if zbc,ok:=m.Load("zbc");ok{
		switch zbc.(type) {
		case map[string]int:
			//zbc["1"]=1
			fmt.Println("okkkkkk")
		default:
			fmt.Println("fuckman")
		}
	}

	if zbc,ok:=m.Load("zbc");ok{
		if p,ok:=zbc.(map[string]int);ok{
			p["zbc"]=1
			fmt.Println(p["zbc"])
		}
		//zbc=zbc.(map[string]int)
		//zbc["zbc"]=1
		//fmt.Println(zbc.(map[string]int))

	}
}