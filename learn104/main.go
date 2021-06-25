package main

import (
	"fmt"
	"sync"
)

type People struct{}

func (p *People) ShowA() {
fmt.Println("showA")
p.ShowB()
}
func (p *People) ShowB() {
fmt.Println("showB")
}

type Teacher struct {
People
}

func (t *Teacher) ShowB() {
fmt.Println("teacher showB")
}

func main() {
	chcat:=make(chan struct{},1)
	chdog:=make(chan struct{},1)
	chflash:=make(chan struct{},1)
	var wg sync.WaitGroup
	wg.Add(100)
	for i:=1;i<=100;i++{
		go func() {
			chcat <- struct{}{}
			fmt.Println("cat!")
		}()
		go func() {
			<-chcat
			fmt.Println("dog!")
			chdog <- struct{}{}
		}()
		go func() {
			<-chdog
			fmt.Println("flash!")
			chflash <- struct{}{}
			wg.Done()
		}()
		<-chflash
	}

	wg.Wait()

}