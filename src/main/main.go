package main

import "trees/trees"

func main() {
	tree := trees.New_splay_tree(3)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(1)
	tree.PrintTree()
}
