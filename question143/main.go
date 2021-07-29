package main

import "fmt"

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

type S struct {
	*T
}

func main() {
	s := S{}
	fmt.Printf("%#v", s) // 输出：main.S{T:(*main.T)(nil)}
	_ = s.foo
	s.foo()
	//_ = s.bar //因为 s.bar 将被展开为 (*s.T).bar，⽽ s.T 是个空指针，解引⽤会 panic。
}
