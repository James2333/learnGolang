package main

import (
	"encoding/json"
	"fmt"
	"learn101/elevator"
	"log"
	"net/http"
)

type work interface {
	Do()
	Run()
}
type workRun struct {
	st []work
}

func NewWorkRun() *workRun  {
	return &workRun{}
}
type Ints struct {
	Name string
}

func NewInts() *Ints {
	return &Ints{Name: "zbc"}
}
func (i *Ints)Do()  {
	fmt.Printf("DO")
}
func (i *Ints)Run()  {
	fmt.Printf("RUN")
}
var els elevator.Elevators
//els:=elevatorTcpClient.NewTestEls()
func main() {
	////接口类型的数组，只要实现这个接口的结构体都可被append进去。
	//sts:=NewWorkRun()
	//subTask1:=NewInts()
	////var s interface{}
	//sts.st=append(sts.st,subTask1)
	////sts.st=append(sts.st,s)
	//for key,value := range sts.st{
	//	fmt.Println(key)
	//	value.Do()
	//	value.Run()
	//}
	//go elevatorTcpClient.NewTcpService()
	//go elevatorTcpClient.ClientSocket()
	//for{
	//	time.Sleep(time.Second)
	//}
	//s:=NewInts()
	//var m sync.Map
	//m.Store("int",*s)
	//fmt.Println(m.Load("int"))
	//test(s)
	//if zbc,ok:=m.Load("int");!ok{
	//	fmt.Println("int ont found")
	//}
	//fmt.Println(zbc)
	//switch v:=zbc.(type) {
	//
	//}
	//if zbc.(type)== Ints {
	//	zbb, ok := zbc.(Ints)
	//	fmt.Println(zbb.Run, ok)

	//str:="qwertyyui"
	//strRune:=[]rune(str)
	//for _,v:=range strRune{
	//	fmt.Printf("%c\n",v)
	//}
	//strRune[1]='Q'
	//str=string(strRune)
	//fmt.Println(str)


	//pool:=&sync.Pool{
	//	New: func() interface{}{
	//	return 0
	//},
	//}
	//// 看一下初始的值，这里是返回0，如果不设置New函数，默认返回nil
	//init := pool.Get()
	//fmt.Println(init)
	//
	//// 设置一个参数1
	//pool.Put(1)
	//
	//// 获取查看结果
	//num := pool.Get()
	//fmt.Println(num)
	//
	//// 再次获取，会发现，已经是空的了，只能返回默认的值。
	//num = pool.Get()
	//fmt.Println(num)

	// 造场景，设置为单核那么就只能是并发，因为go1.5版本之后，默认是多核了。
	//runtime.GOMAXPROCS(1)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println("go")
	//	}
	//}()
	//
	//for i := 0; i < 2; i++ {
	//	runtime.Gosched()
	//	fmt.Println("hello")
	//}


	//ch := make(chan struct{})
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	close(ch)
	//}()
	//
	//fmt.Println("脑子好像进...")
	//<-ch
	//fmt.Println("煎鱼了！")

	//opel:=elevator.NewTestOperationEl()
	////opelbyte:=make([]byte,100)
	//opelbytes,err:=json.Marshal(opel)
	//if err != nil {
	//	fmt.Printf("Marshal falid %v",err)
	//}
	//
	//a:=&elevator.OperationEl{}
	//err=json.Unmarshal(opelbytes,a)
	//if err != nil {
	//	fmt.Printf("Unmarshal faild %v",err)
	//}
	//fmt.Println(a)

	//for _,v:=range els{
	//	log.Printf("%v",v)
	//}
	http.HandleFunc("/",GetRightElevator)
	http.ListenAndServe(":8000",nil)

}

func GetRightElevator(w http.ResponseWriter,r *http.Request)  {
	els:=NewTestEls()
	var result struct {
		eleID string
	}
	for _,v:=range els{
		log.Printf("改变值之前%v",v)
	}
	result.eleID,_=els.RightElevator(-1,3)
	for _,v:=range els{
		log.Printf("改变值之后%v",v)
	}
	resJson,_:=json.Marshal(result)
	w.Write(resJson)
}


func NewTestEls() elevator.Elevators {
	els := elevator.NewElevators()
	el1 := &elevator.Elevator{
		ElevatorId:   "1",
		Floor:        3,
		State:        "0",
		CurrentState: "1",
	}
	el2 := &elevator.Elevator{
		ElevatorId:   "2",
		Floor:        2,
		State:        "0",
		CurrentState: "1",
	}
	el3 := &elevator.Elevator{
		ElevatorId:   "3",
		Floor:        1,
		State:        "0",
		CurrentState: "0",
	}
	el4 := &elevator.Elevator{
		ElevatorId:   "4",
		Floor:        -1,
		State:        "0",
		CurrentState: "0",
	}
	els.Update(el1)
	els.Update(el2)
	els.Update(el3)
	els.Update(el4)
	return els
}
func test(value interface{})  {
	switch v:=value.(type) {
	case Ints:
		op,_:=value.(Ints)
		fmt.Println(op.Name)
	default:
		fmt.Println(v)
	}

}