package main

import "fmt"

func main() {
	var m map[int]bool // nil
	_ = m[123]
	var p *[5]string // nil
	fmt.Printf("%#v\n", p)
	for range p {
		_ = len(p)
		fmt.Println(len(p))
	}
	var s []int // nil
	_ = s[:]
	s, s[0] = []int{1, 2}, 9 //panic: runtime error: index out of range [0] with length 0
	//因为左侧的 s[0] 中的 s 为 nil。
}
