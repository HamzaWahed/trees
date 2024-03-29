package main

import (
	"fmt"
	"trees/trees"
)

func main() {
	tree := trees.BuildVEB(4)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	fmt.Println(tree.Summary.Min)
	fmt.Println(tree.Summary.Max)
	fmt.Println(tree.Predecessor(0))

}
