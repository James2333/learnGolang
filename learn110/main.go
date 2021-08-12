package main

import "fmt"

// 父接口
type Humen interface {
	SayHello()
}

// 子接口
type Person interface {
	Humen  // 继承父接口
	sing(string)
}

// 学生类
type Student struct {
	name string
	age int
}

type Fuckman struct {
	Humen
	Person
}

// 学生类的方法 (让学生类符合父接口的规则)
func (stu *Student) SayHello() {
	fmt.Printf("我是学生，名字是%s，我%d岁了\n", stu.name, stu.age)
}
// 学生类的方法 (让学生类符合子接口的规则)
func (stu *Student) sing(str string) {
	fmt.Printf("唱歌：%s\n", str)
}


func main() {
	// 声明接口类型的变量
	var h Humen  // 父接口
	var per Person  // 子接口

	// 学生对象
	stu := Student{"张三", 20}
	stu.SayHello()

	per = &stu  // 子接口
	per.SayHello()
	per.sing("啦啦啦。。。")

	h = per  // 父接口
	h.SayHello()
}

