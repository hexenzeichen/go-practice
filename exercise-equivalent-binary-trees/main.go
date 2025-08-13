package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Walking(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	defer close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	x := make(chan int)
	y := make(chan int)

	go Walking(t1, x)
	go Walking(t2, y)

	for {
		v1, ok1 := <-x
		v2, ok2 := <-y

		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break;
		}
	}

	return true
}

func main() {
	if Same(tree.New(1), tree.New(1)) {
		fmt.Print("Yes!")
	} else {
		fmt.Print("No!")
	}
}
