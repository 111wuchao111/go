package main

import "fmt"

type T struct {
	name string
}

var t = T{"wu"}

var y = h()

func main() {
	fmt.Println(t)
	fmt.Println(y)
}

func h() int {
	return 6
}
