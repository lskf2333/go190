package main

import "fmt"

//这道题会不会有点诧异，我们分配给变量 p 的值明明是 nil，然⽽ p 却不是 nil。记住⼀点，当且仅当动
//态值和动态类型都为 nil 时，接⼝类型值才为 nil。上⾯的代码，给变量 p 赋值之后，p 的动态值是
//nil，但是动态类型却是 *Student，是⼀个 nil 指针，所以相等条件不成⽴。

type People interface {
	Show()
}
type Student struct{}

func (stu *Student) Show() {
}
func main() {
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
