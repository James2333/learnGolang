package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func busi(ch chan bool, i int) {

	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
}

func main() {
	//模拟用户需求业务的数量
	//task_cnt := math.MaxInt64
	////task_cnt := 10
	//
	//ch := make(chan bool, 3)
	//
	//for i := 0; i < task_cnt; i++ {
	//
	//	ch <- true
	//
	//	go busi(ch, i)
	//}
	var m sync.RWMutex
	for i:=0;i<10;i++{
		go func(i int) {
			m.RLock()
			fmt.Println("RWMutex read lock ",i)
			defer m.RUnlock()
			//var t time.Duration
			t:=time.Second*time.Duration(i)
			time.Sleep(t)
			fmt.Println("RWMutex read Ulock",i)
		}(i)
	}
	go func() {
		m.Lock()
		fmt.Println("RWMutex write lock ")
		defer m.Unlock()
		time.Sleep(time.Second*5)
		fmt.Println("RWMutex write Ulock")
	}()
	//time.Sleep("")
	for  {
		time.Sleep(time.Second)
	}
}