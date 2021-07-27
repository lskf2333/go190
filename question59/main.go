package main

import "fmt"

type myInt int

func (i myInt) PrintInt () {
	fmt.Println(i)
}

//方法必须和类型在同一个包，int 不在main 包
//基于类型创建的⽅法必须定义在同⼀个包内，上⾯的代码基于 int 类型创建了 PrintInt() ⽅法，由于 int 类型和⽅法 PrintInt() 定义在不同的包内，所以编译出错。
//func (i int) PrintInt1 () {
//	fmt.Println(i)
//}

func main() {
	var i myInt = 1
	i.PrintInt()

	//var j int =2
	//j.PrintInt1()
}