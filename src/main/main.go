package main

import "trees/trees"

func main() {
	tree := trees.NewTangoTree([]int{2, 4, 8, 6, 12, 10, 14})
	trees.PrintAuxNodes(tree.Root)
}
