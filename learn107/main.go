package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"learn101/elevator"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

type work interface {
	Do()
	Run()
}
type workRun struct {
	st []work
}

func NewWorkRun() *workRun {
	return &workRun{}
}

type Ints struct {
	Name string
}

func NewInts() *Ints {
	return &Ints{Name: "zbc"}
}
func (i *Ints) Do() {
	fmt.Printf("DO")
}
func (i *Ints) Run() {
	fmt.Printf("RUN")
}

var els elevator.Elevators

//els:=elevatorTcpClient.NewTestEls()

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func iplus() func() int {
	i := 1
	defer func() { i++ }()
	return func() int {
		i++
		return i
	}

}

type Package struct {
	Version        [2]byte // 协议版本
	Length         int16   // 数据部分长度
	Timestamp      int64   // 时间戳
	HostnameLength int16   // 主机名长度
	Hostname       []byte  // 主机名
	TagLength      int16   // Tag长度
	Tag            []byte  // Tag
	Msg            []byte  // 数据部分长度
}

func (p *Package) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.Version)
	err = binary.Write(writer, binary.BigEndian, &p.Length)
	err = binary.Write(writer, binary.BigEndian, &p.Timestamp)
	err = binary.Write(writer, binary.BigEndian, &p.HostnameLength)
	err = binary.Write(writer, binary.BigEndian, &p.Hostname)
	err = binary.Write(writer, binary.BigEndian, &p.TagLength)
	err = binary.Write(writer, binary.BigEndian, &p.Tag)
	err = binary.Write(writer, binary.BigEndian, &p.Msg)
	return err
}
func (p *Package) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.Version)
	err = binary.Read(reader, binary.BigEndian, &p.Length)
	err = binary.Read(reader, binary.BigEndian, &p.Timestamp)
	err = binary.Read(reader, binary.BigEndian, &p.HostnameLength)
	p.Hostname = make([]byte, p.HostnameLength)
	err = binary.Read(reader, binary.BigEndian, &p.Hostname)
	err = binary.Read(reader, binary.BigEndian, &p.TagLength)
	p.Tag = make([]byte, p.TagLength)
	err = binary.Read(reader, binary.BigEndian, &p.Tag)
	p.Msg = make([]byte, p.Length-8-2-p.HostnameLength-2-p.TagLength)
	err = binary.Read(reader, binary.BigEndian, &p.Msg)
	return err
}

func (p *Package) String() string {
	return fmt.Sprintf("version:%s length:%d timestamp:%d hostname:%s tag:%s msg:%s",
		p.Version,
		p.Length,
		p.Timestamp,
		p.Hostname,
		p.Tag,
		p.Msg,
	)
}
func CreateContainer(w http.ResponseWriter,r *http.Request)  {
	container := make([]byte,1024*1024*100)
	for i:=0;i<len(container);i++{
		container[i]=0
	}

}
func main() {


	runtime.GOMAXPROCS(1)
	go func(){
		fmt.Println("hello world")
		// panic("hello world")  // 强制观察输出
	}()
	go func(){
		for {
			time.Sleep(time.Second)
			//fmt.Println("aaa")  // 非内联函数，这行注释打开，将导致 hello world 的输出
		}
	}()
	select {}

	//sumfunc:=make([]func(),10)
	//sumfunc=append(sumfunc, func() {
	//	fmt.Println("zbc")
	//})
	//for k,v:=range sumfunc{
	//	fmt.Println(k,v)
	//}
	//sumfunc[10]()
	//chans:=make(chan int,10)
	//sum:=make([]int,10)
	//for i:=range sum{
	//	sum[i]=i
	//}
	//reverce(sum,chans)
	//for i:=range sum{
	//	sum[i]=<-chans
	//}
	//for i:=range sum{
	//	fmt.Println(sum[i])
	//}
	//http.Handle("/",CreateContainer)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	//var i interface{}
	//i = &sli
	//fmt.Println(i)
	//sli[0]=10
	//fmt.Println(i)
	//v := []int{1, 2, 3}
	//for i := range v {
	//	v = append(v, i)
	//}
	//for _,v := range v {
	//	fmt.Println(v)
	//}
	//<-time.After(time.Second)
	//hostname, err := os.Hostname()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//pack := &Package{
	//	Version:        [2]byte{'V', '1'},
	//	Timestamp:      time.Now().Unix(),
	//	HostnameLength: int16(len(hostname)),
	//	Hostname:       []byte(hostname),
	//	TagLength:      4,
	//	Tag:            []byte("demo"),
	//	Msg:            []byte(("现在时间是:" + time.Now().Format("2006-01-02 15:04:05"))),
	//}
	//pack.Length = 8 + 2 + pack.HostnameLength + 2 + pack.TagLength + int16(len(pack.Msg))
	//
	//buf := new(bytes.Buffer)
	//// 写入四次，模拟TCP粘包效果
	//pack.Pack(buf)
	//pack.Pack(buf)
	//pack.Pack(buf)
	//pack.Pack(buf)
	//// scanner
	//scanner := bufio.NewScanner(buf)
	//scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	//	if !atEOF && data[0] == 'V' {
	//		if len(data) > 4 {
	//			length := int16(0)
	//			binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &length)
	//			if int(length)+4 <= len(data) {
	//				return int(length) + 4, data[:int(length)+4], nil
	//			}
	//		}
	//	}
	//	return
	//})
	//for scanner.Scan() {
	//	scannedPack := new(Package)
	//	scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
	//	log.Println(scannedPack)
	//}
	//if err := scanner.Err(); err != nil {
	//	log.Fatal("无效数据包")
	//}
}

func reverce(s []int,chans chan int) []int {
	for i:=range s{
		defer func(i int)  {
			chans<-i
		}(i)

	}
	return s
}
//func main() {
//	////接口类型的数组，只要实现这个接口的结构体都可被append进去。
//	//sts:=NewWorkRun()
//	//subTask1:=NewInts()
//	////var s interface{}
//	//sts.st=append(sts.st,subTask1)
//	////sts.st=append(sts.st,s)
//	//for key,value := range sts.st{
//	//	fmt.Println(key)
//	//	value.Do()
//	//	value.Run()
//	//}
//	//go elevatorTcpClient.NewTcpService()
//	//go elevatorTcpClient.ClientSocket()
//	//for{
//	//	time.Sleep(time.Second)
//	//}
//	//s:=NewInts()
//	//var m sync.Map
//	//m.Store("int",*s)
//	//fmt.Println(m.Load("int"))
//	//test(s)
//	//if zbc,ok:=m.Load("int");!ok{
//	//	fmt.Println("int ont found")
//	//}
//	//fmt.Println(zbc)
//	//switch v:=zbc.(type) {
//	//
//	//}
//	//if zbc.(type)== Ints {
//	//	zbb, ok := zbc.(Ints)
//	//	fmt.Println(zbb.Run, ok)
//
//	//str:="qwertyyui"
//	//strRune:=[]rune(str)
//	//for _,v:=range strRune{
//	//	fmt.Printf("%c\n",v)
//	//}
//	//strRune[1]='Q'
//	//str=string(strRune)
//	//fmt.Println(str)
//
//	//pool:=&sync.Pool{
//	//	New: func() interface{}{
//	//	return 0
//	//},
//	//}
//	//// 看一下初始的值，这里是返回0，如果不设置New函数，默认返回nil
//	//init := pool.Get()
//	//fmt.Println(init)
//	//
//	//// 设置一个参数1
//	//pool.Put(1)
//	//
//	//// 获取查看结果
//	//num := pool.Get()
//	//fmt.Println(num)
//	//
//	//// 再次获取，会发现，已经是空的了，只能返回默认的值。
//	//num = pool.Get()
//	//fmt.Println(num)
//
//	// 造场景，设置为单核那么就只能是并发，因为go1.5版本之后，默认是多核了。
//	//runtime.GOMAXPROCS(1)
//	//go func() {
//	//	for i := 0; i < 5; i++ {
//	//		fmt.Println("go")
//	//	}
//	//}()
//	//
//	//for i := 0; i < 2; i++ {
//	//	runtime.Gosched()
//	//	fmt.Println("hello")
//	//}
//
//	//ch := make(chan struct{})
//	//go func() {
//	//	time.Sleep(1 * time.Second)
//	//	close(ch)
//	//}()
//	//
//	//fmt.Println("脑子好像进...")
//	//<-ch
//	//fmt.Println("煎鱼了！")
//
//	//opel:=elevator.NewTestOperationEl()
//	////opelbyte:=make([]byte,100)
//	//opelbytes,err:=json.Marshal(opel)
//	//if err != nil {
//	//	fmt.Printf("Marshal falid %v",err)
//	//}
//	//
//	//a:=&elevator.OperationEl{}
//	//err=json.Unmarshal(opelbytes,a)
//	//if err != nil {
//	//	fmt.Printf("Unmarshal faild %v",err)
//	//}
//	//fmt.Println(a)
//
//	//for _,v:=range els{
//	//	log.Printf("%v",v)
//	//}
//	//http.HandleFunc("/",GetRightElevator)
//	//http.ListenAndServe(":8000",nil)
//	//zbc:=iplus()
//	//fmt.Println(zbc())
//	//fmt.Println(zbc())
//	//fmt.Println(zbc())
//	//
//	//	nextInt := intSeq()
//	//
//	//	fmt.Println(nextInt())
//	//	fmt.Println(nextInt())
//	//	fmt.Println(nextInt())
//	//
//	//	newInts := intSeq()
//	//	fmt.Println(newInts())
//
//	//messages := make(chan string)
//	//signals := make(chan bool)
//	//
//	//select {
//	//case msg := <-messages:
//	//	fmt.Println("received message", msg)
//	//default:
//	//	fmt.Println("no message received")
//	//}
//	//
//	//msg := "hi"
//	//select {
//	//case messages <- msg:
//	//	fmt.Println("sent message", msg)
//	//default:
//	//	fmt.Println("no message sent")
//	//}
//	//
//	//select {
//	//case msg := <-messages:
//	//	fmt.Println("received message", msg)
//	//case sig := <-signals:
//	//	fmt.Println("received signal", sig)
//	//default:
//	//	fmt.Println("no activity")
//	//}
//
//	//timer := time.NewTimer(3 * time.Second)
//	//for {
//	//	timer.Reset(3 * time.Second) // 这里复用了 timer
//	//	select {
//	//	case <-timer.C:
//	//		fmt.Println("每隔3秒执行一次")
//	//	}
//	//}
//
//	//requests := make(chan int, 5)
//	//for i := 1; i <= 5; i++ {
//	//	requests <- i
//	//}
//	//close(requests)
//	//
//	//limiter := time.Tick(2000 * time.Millisecond)
//	//
//	//for req := range requests {
//	//	<-limiter
//	//	fmt.Println("request", req, time.Now())
//	//}
//	//
//	//burstyLimiter := make(chan time.Time, 3)
//	//
//	//for i := 0; i < 3; i++ {
//	//	burstyLimiter <- time.Now()
//	//}
//	//
//	//go func() {
//	//	for t := range time.Tick(200 * time.Millisecond) {
//	//		burstyLimiter <- t
//	//	}
//	//}()
//	//
//	//burstyRequests := make(chan int, 5)
//	//for i := 1; i <= 5; i++ {
//	//	burstyRequests <- i
//	//}
//	//close(burstyRequests)
//	//for req := range burstyRequests {
//	//	<-burstyLimiter
//	//	fmt.Println("request", req, time.Now())
//	//}
//	//a:=make([]int,2)
//	//b:=make([]int,2)
//	//a=append(a,b...)
//	//fmt.Println(a)
//
//	//a := Fun()
//	//b:=a("hello ")
//	//c:=a("hello ")
//	//fmt.Println(b)//worldhello
//	//fmt.Println(c)//worldhello hello
//	////fmt.Println(c)//worldhello hello
//	////fmt.Println(c)//worldhello hello
//	//zbc:=Addint()
//	//fmt.Println(zbc())
//	//fmt.Println(zbc())
//	//fmt.Println(zbc())
//
//	//go spinner(100 * time.Millisecond)
//	//const n = 45
//	//fibN := fib(n) // slow
//	//fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
//
//	//var once sync.Once
//	//onceBody := func() {
//	//	fmt.Println("Only once")
//	//}
//	//done := make(chan bool)
//	//for i := 0; i < 10; i++ {
//	//	go func() {
//	//		once.Do(onceBody)
//	//		done <- true
//	//	}()
//	//}
//	//for i := 0; i < 10; i++ {
//	//	fmt.Printf("%v\n",<-done)
//	//}
//
//	//var studentPool = sync.Pool{
//	//	New: func() interface{} {
//	//		return new(Student)
//	//	},
//	//}
//	//stu := studentPool.Get().(*Student)
//	//
//	//json.Unmarshal(buf, stu)
//	//fmt.Printf("%v",stu)
//	//studentPool.Put(stu)
//
//	//var num float64 = 1.2345
//	//
//	//pointer := reflect.ValueOf(&num)
//	//value := reflect.ValueOf(num)
//	//
//	//// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
//	//// Golang 对类型要求非常严格，类型一定要完全符合
//	//// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
//	//convertPointer := pointer.Interface().(*float64)
//	//convertValue := value.Interface().(float64)
//	//
//	//fmt.Println(convertPointer)
//	//fmt.Println(convertValue)
//	zbc:=make(chan interface{},10)
//	quit:=make(chan struct{})
//	go func() {
//		time.Sleep(time.Second*15)
//		quit<- struct{}{}
//	}()
//	go func() {
//		user := User{1, "Allen.Wu", 25}
//		student:=Student{
//			Name:   "fuckman",
//			Age:    18,
//			Remark: [1024]byte{},
//		}
//		for  {
//			zbc<-user
//			time.Sleep(time.Second*3)
//			zbc<-student
//			time.Sleep(time.Second*10)
//		}
//	}()
//
//	for  {
//		select {
//		case zbb:=<-zbc:
//			DoFiledAndMethod(zbb)
//		case <-quit:
//			return
//		}
//	}
//
//
//}


type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}


type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func unmarsh() {
	stu := &Student{}
	json.Unmarshal(buf, stu)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
func Addint() func()int{
	a:=0
	return func() int {
		a++
		return a
	}
}
func Fun() func(string) string {
	a := "world"
	return func(args string) string {
		a += args
		return  a
	}
}
func GetRightElevator(w http.ResponseWriter, r *http.Request) {
	els := NewTestEls()
	var result struct {
		eleID string
	}
	for _, v := range els {
		log.Printf("改变值之前%v", v)
	}
	result.eleID, _ = els.RightElevator(-1)
	for _, v := range els {
		log.Printf("改变值之后%v", v)
	}
	resJson, _ := json.Marshal(result)
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
func test(value interface{}) {
	switch v := value.(type) {
	case Ints:
		op, _ := value.(Ints)
		fmt.Println(op.Name)
	default:
		fmt.Println(v)
	}

}
