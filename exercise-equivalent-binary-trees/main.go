package main

import "golang.org/x/tour/tree"

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

func Same(t1, t2 *tree.Tree) bool

func main() {
}
