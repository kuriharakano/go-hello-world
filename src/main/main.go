package main

import (
	. "fmt"
	_ "github.com/zooh82/go-hello-world/src/gosample"
	_ "gosample"
)

var message string = "hello world"
var (
	a string = "aaa"
	b        = "bbb"
	c        = "ccc"
	d        = "ddd"
)

const msg string = "This message is a constant."

func calc() {
	// i := 0
	var j int
	immsg := "Implicit conversion"
	Println(immsg)
	Println(j)

	a, b := 1000, 100

	if a < b {
		Println("b bigger than a")
	} else if b < a {
		Println("a bigger than b")
	} else {
		Println("a equals b")
	}

	for i := 0; i < 10; i++ {
		Println(i)
	}

	n := 3
	switch n {
	case 3:
		n = n - 1
		fallthrough
	case 2:
		n = n - 1
		fallthrough
	case 1:
		n = n - 1
		Println(n) // 0
	}
}

// main function
func main() {
	// Println(sum(1, 2))
	// sayHello()
	x, y := 3, 4

	x, y = swap(x, y)
	Println(x)
	Println(y)
}

func swap(i, j int) (int, int) {
	return j, i
}

func sum(i, j int) int {
	return i + j
}
func sayHello() {
	Println("hello !")
}
