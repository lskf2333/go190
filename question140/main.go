package main

import "fmt"

func main() {
	x := make([]int, 2, 10)
	y := x[6:10]
	//_ = x[6:] 截取符号 [i:j]，如果 j 省略，默认是原切⽚或者数组的⻓度，x 的⻓度是 2，⼩于起始下标 6 ， 所以 panic。
	z := x[2:]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
