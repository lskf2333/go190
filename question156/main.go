package main

import "fmt"

//关联question155,157,158

type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ") //9
		}
	}
	fmt.Print(ts) //[{0} {9}]
}

//知识点：for-range 数组指针。
//for-range 循环中的循环变量 t 是原数组元素的副本。如果数组元素是结构体值，则副本的字段和原数组字段是两个不同的值。
