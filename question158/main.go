package main

import "fmt"

//关联question155,156,157

type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}

//知识点：for-range 切⽚。
//for-range 切⽚时使⽤的是切⽚的副本，但不会复制底层数组，换句话说，此副本切⽚与原数组共享底层数组。
