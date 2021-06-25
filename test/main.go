package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type WorkPool struct {
	worklen int

}
type Limit struct {
	channel chan struct{}
}

func NewLimit(number int) *Limit {
	return &Limit{
		channel: make(chan struct{}, number),
	}
}
func NewQueue(number int) *Limit {
	return &Limit{
		channel: make(chan struct{}, number),
	}
}
func WriteQueue(queue *Limit,limit *Limit)  {
	for {
		select {
		case <-queue.channel:
			{
				limit.channel <- struct{}{}
				//全局队列往执行队列里面输入任务
			}
		}
	}
}
func (limit *Limit) Run(wg *sync.WaitGroup) {
	go func() {
		//模拟程序处理时间，随机生成一百以内的整数去乘以一毫秒
		rand.Seed(time.Now().Unix())
		slp:=rand.Intn(1000)
		zbc:=rand.Intn(8)
		time.Sleep(time.Millisecond*time.Duration(slp))
		for i:=1;i<=zbc;i++{
			fmt.Printf("-")
		}
		fmt.Println()
		wg.Done()
		<-limit.channel
	}()
}

const number = 100000

type Count struct {
	a int
	b int
}

func NewCount(a,b int) *Count {
	return &Count{
		a: a,
		b: b,
	}
}
func (c *Count) Counts()*Count {
	return &Count{
		a: c.a+c.b,
		b: c.a,
	}
}

func product(msg chan interface{})  {
	for  {
		rand.Seed(time.Now().Unix())
		msg<-rand.Float64()
		time.Sleep(time.Second/10)
	}
}
func consumer(msg chan interface{})  {
	for  {
		read:=<-msg
		fmt.Println(read)
	}

}
func main() {
	msg:=make(chan interface{})
	go product(msg)
	go product(msg)
	go product(msg)
	go product(msg)
	go product(msg)
	go consumer(msg)
	go consumer(msg)
	consumer(msg)
	//count:=NewCount(2,3)
	//defer fmt.Println(count.Counts().Counts().Counts())
	//Time := time.Now()
	//wg := sync.WaitGroup{}
	//wg.Add(number)
	////初始化全局队列
	////queue := NewQueue(number)
	////初始化并发队列
	//limit := NewLimit(number)
	////模拟http请求
	//go func(){
	//	for i := 1; i <= number; i++ {
	//		//queue.channel <- struct{}{}
	//		limit.channel<- struct{}{}
	//		limit.Run(&wg)
	//		//limit.Run(wg)
	//	}
	//}()
	////for  {
	////	select {
	////	case <-queue.channel:
	////		limit.channel<- struct{}{}
	////		go limit.Run(&wg)
	////	default:
	////		break
	////	}
	////}
	////go func() {
	////	WriteQueue(queue,limit)
	////}()
	////从执行队列中获取任务执行，调用goroutine
	//wg.Wait()
	//for {
	//	fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
	//	time.Sleep(2 * time.Second)
	//	fmt.Println(time.Since(Time))
	//	fmt.Println("EXIT!!!!!!!!!!!!!!!")
	//}
}
