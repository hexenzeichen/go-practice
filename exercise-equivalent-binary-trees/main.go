package main

import "golang.org/x/tour/tree"

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}



func Walk(t *tree.Tree, ch chan int)

func Same(t1, t2 *tree.Tree) bool

func main() {
}
