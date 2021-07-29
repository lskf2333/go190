package main

import "fmt"

//关联question152

func main() {
	defer func() {
		fmt.Print(recover()) //2
	}()
	defer func() {
		defer func() {
			fmt.Print(recover()) //1
		}()
		panic(1)
	}()
	defer recover()
	panic(2)
}
