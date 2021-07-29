package main

func test(x int) (func(), func()) {
	println(&x)
	return func() {
			println(x)
			println(&x)
			x += 10
		}, func() {
			println(x)
			println(&x)
		}
}
func main() {
	a, b := test(100)
	a() //100
	b() //110
}

//知识点：闭包引⽤相同变量。
