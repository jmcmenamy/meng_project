package main

import (
	"fmt"
)

type data struct {
	num	int64
	comment	string
	elems	map[int]bool
}

func Debug(msg string) {
	fmt.Printf("%v\n", msg)
}

func (d *data) foo(x int, g string) int {
	y := 1
	fmt.Printf("Variables seen so far: %v %v %v\n", x, g, y)
	fmt.Printf("Variables seen so far: %v %v %v\n", x, g, y)
	Debug("testing this!")
	return x + y
}
