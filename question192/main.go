package main

import (
	"fmt"
)

func main() {
	a := [3]int{0, 1, 2}
	s := a[1:2]
	fmt.Println(a)
	fmt.Println(s)
	fmt.Printf("", &s)
	fmt.Println(&s)
	fmt.Println(`!!!!!`)
	s[0] = 11
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(&s)
	fmt.Println(&s)
	fmt.Println(`!!!!!`)
	s = append(s, 12)
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(&s)
	fmt.Println(&s)
	fmt.Println(`!!!!!`)
	s = append(s, 13)
	s[0] = 21
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(&s)
	fmt.Println(&s)
	fmt.Println(`!!!!!`)
}
