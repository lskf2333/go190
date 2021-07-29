package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

//知识点：WaitGroup 的使⽤。存在两个问题：
//在协程中使⽤ wg.Add()；
//使⽤了 sync.WaitGroup 副本；
