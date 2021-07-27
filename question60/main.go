package main

import "fmt"

type People interface {
	Speak(string) string
}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	//编译错误 Student does not implement People (Speak method has pointer receiver) ，
	//值类型 Student 没有实现接⼝的 Speak() ⽅法，⽽是指针类型 *Student 实现该⽅法。
	//var peo People = Student{}
	var peo People =&Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
