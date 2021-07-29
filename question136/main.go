package main

//关联question141
//当指针值赋值给变量或者作为函数参数传递时，会⽴即计算并复制该⽅法执⾏所需的接收者对象，与其
//绑定，以便在稍后执⾏时，能隐式第传⼊接收者参数。
import "fmt"

type N int

func (n N) test() {
	fmt.Println(n)
}
func main() {
	var n N = 10
	p := &n
	n++
	f1 := n.test
	n++
	f2 := p.test
	n++
	fmt.Println(n)
	f1()
	f2()
}
