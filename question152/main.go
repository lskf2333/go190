package main

import "fmt"

//关联question154

func main() {
	defer func() {
		fmt.Print(recover()) //1
	}()
	defer func() {
		defer fmt.Print(recover()) //2
		panic(1)
	}()
	defer recover()
	panic(2)
}

//recover() 必须在 defer() 函数中调⽤才有效，所以第 9 ⾏代码捕获是⽆效的。在调⽤ defer() 时，便会计
//算函数的参数并压⼊栈中，所以执⾏第 6 ⾏代码时，此时便会捕获 panic(2)；此后的 panic(1)，会被上
//⼀层的 recover() 捕获。所以输出 21。
