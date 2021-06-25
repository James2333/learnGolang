package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

func isqueString(str string) bool {
	if strings.Count(str, "") > 3000 {
		return false
	}
	for _, v := range str {
		if v > 127 {
			return false
		}
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true

}
func reverString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	if len(str) > 5000 {
		return s, false
	}
	for i := 0; i < len(str)/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str), true
}
func isRegroup(s1, s2 string) bool {
	str1 := []rune(s1)
	str2 := []rune(s2)
	if len(str1) > 5000 || len(str2) > 5000 || len(str1) != len(str2) {
		return false
	}
	for _, v := range str1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true

}

type param map[string]interface{}
type zbc struct {
	param
}
type People struct {
	Name string `json:"name"`
}

func (p *People) String() string {
	return fmt.Sprintf("print:%v", p)

}

func zbb(i int) {
	fmt.Println("zbc:", i)
}

type UserAge struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAge) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age

}
func sb() <-chan interface{} {
	ch := make(chan interface{})
	close(ch)
	return ch
}

var rw sync.RWMutex

func (ua *UserAge) Get(name string) int {
	rw.RLock()
	defer rw.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1

}

type Peoples interface {
	show()
}
type student struct {
}

//func (stu *student)speak(think string)(talk string)  {
//	if think=="bitch"{
//		talk="you are good boy!"
//	}else {
//		talk="zbc"
//	}
//	return
//}
//func (stu *student)Show()  {
//
//}
//func live() Peoples {
//	var peo *student
//	return peo
//}
type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}
type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return

}

type Ban struct {
	visitIP map[string]time.Time
	lock    sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{
		visitIP: make(map[string]time.Time),
	}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				o.lock.Lock()
				for k, v := range o.visitIP {
					if time.Now().Sub(v) >= time.Minute*1 {
						delete(o.visitIP, k)
					}
				}
				o.lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return

			}
		}
	}()
	return o
}
func proc()  {
	panic("ok")
}
func main() {
	go func() {
		defer func() {
			if err:=recover();err!=nil{
				fmt.Println(err)
		}
		}()
		proc()
	}()
	time.Sleep(time.Second/2)
	//go func() {
	//	t:=time.NewTicker(time.Second*1)
	//	for{
	//		select {
	//		case <-t.C:
	//				go func() {
	//					defer func() {
	//						if err:=recover();err!=nil{
	//							fmt.Println(err)
	//						}
	//					}()
	//					proc()
	//				}()
	//		}
	//	}
	//}()
	//select {
	//
	//}
	//ch:=make(chan int,5)
	//for i:=0;i<5;i++{
	//	go func() {
	//		ch<-rand.Intn(1000)
	//	}()
	//	select {
	//	case number:=<-ch:
	//		fmt.Println(number)
	//	}
	//}
	//time.Sleep(time.Second*2)
	//if live()==nil{
	//	fmt.Println("asdasdad")
	//}else {
	//	fmt.Println("zxcvbnm")
	//}
	//var peo Peoples =&student{}
	//think:="bitch"
	//fmt.Println(peo.speak(think))
	//ch:=sb()

	//zbb:=&UserAge{}
	//zbb.Add("abc",1)
	//fmt.Println(zbb.Get("abc"))

	//i:=1
	//defer zbb(i)
	//i++
	//defer zbb(i)
	//runtime.GOMAXPROCS(1)
	//wg:=sync.WaitGroup{}
	//wg.Add(20)
	//for i:=0;i<10;i++{
	//	go func() {
	//		fmt.Println("i1:",i)
	//		wg.Done()
	//	}()
	//}
	//for i:=0;i<10;i++{
	//	go func(i int) {
	//		fmt.Println("i2:",i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()

	//p:=&People{}
	//p.String()
	//js:=`{"name": "11"}`
	//var p People
	//err:=json.Unmarshal([]byte(js),&p)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(p.Name)
	//s:=&zbc{}
	//s.param["zbc"]=1111
	//fmt.Println(isqueString("abcdefg"))
	//fmt.Println(reverString("zbc"))
	//fmt.Println(isRegroup("zxcvbnm","mnbvcx z1"))
	//s:="1            "
	//fmt.Println(len([]rune(s)))
	//s=strings.Replace(s," ","%20",-1)
	//fmt.Println(s)
}
