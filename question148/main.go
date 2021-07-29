package main

import "fmt"

func main() {
	var x int8 = -128
	var y = x / -1
	fmt.Println(y) //-128  //溢出，int8 的范围：-128(-2^7)~127(2^7-1)
	var x1 int8 = -127
	var y1 = x1 / -1
	fmt.Println(y1) //127
	var x2 int8 = -128
	var y2 = uint8(x2 / -1)
	fmt.Println(y2) //128
}
