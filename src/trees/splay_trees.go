package trees

import (
	"fmt"
)

type Node struct {
	Data       int32
	LeftChild  *Node
	RightChild *Node
}

type SplayTree struct {
	Root *Node
	Size int32
}

const NullPointerError string = "Cannot call search on a null splay tree pointer"
const MemoryAllocationError string = "Could not allocate memory"

// NewSplayTree Initializes a new splay tree node and returns a pointer to it.
func NewSplayTree(x int32) *SplayTree {
	tree := new(SplayTree)
	if tree == nil {
		panic(MemoryAllocationError)
	}

	tree.Root = new(Node)
	if tree.Root == nil {
		panic(MemoryAllocationError)
	}

	tree.Size = 0
	tree.Root.Data = x
	tree.Root.LeftChild = nil
	tree.Root.RightChild = nil
	return tree
}

// Search Binary searches the splay tree for a node with x in its data field. Returns a pointer to the node or to the parent node
// if the x is not in the splay tree. Splays the node to the top if it exists, otherwise splays the node where the search stops on.
func (tree *SplayTree) Search(x int32) bool {
	if tree == nil {
		panic(NullPointerError)
	}

	node := searchHelper(x, tree.Root)
	tree.splay(node)
	if node.Data == x {
		return true
	}

	return false
}

func searchHelper(x int32, node *Node) *Node {
	if node.Data == x {
		return node
	}

	if node.LeftChild == nil && node.RightChild == nil {
		return node
	}

	if node.Data < x {
		if node.RightChild == nil {
			return node
		}

		return searchHelper(x, node.RightChild)
	} else {
		if node.LeftChild == nil {
			return node
		}

		return searchHelper(x, node.LeftChild)
	}
}

// Insert Creates a new splay tree node with x and inserts the node into the tree. Splays the node to the root of
// the tree after insertion.
func (tree *SplayTree) Insert(x int32) bool {
	if tree == nil {
		panic(NullPointerError)
	}

	node := new(Node)
	if node == nil {
		panic(MemoryAllocationError)
	}

	node.Data = x
	insertHelper(tree.Root, node)
	return true
}

func insertHelper(root *Node, node *Node) {
	if root.Data > node.Data {
		if root.LeftChild == nil {
			root.LeftChild = node
		} else {
			insertHelper(root.LeftChild, node)
		}
	} else {
		if root.RightChild == nil {
			root.RightChild = node
		} else {
			insertHelper(root.RightChild, node)
		}
	}
}

func findParent(x int32, node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.LeftChild == nil && node.RightChild == nil {
		return nil
	}

	if node.Data < x {
		if node.RightChild != nil {
			if node.RightChild.Data == x {
				return node
			} else {
				return findParent(x, node.RightChild)
			}
		}
	} else {
		if node.LeftChild != nil {
			if node.LeftChild.Data == x {
				return node
			} else {
				return findParent(x, node.LeftChild)
			}
		}
	}

	return nil
}

func (tree *SplayTree) splay(node *Node) {
	y := findParent(node.Data, tree.Root)
	var z *Node = nil
	var w *Node = nil
	if y != nil {
		z = findParent(y.Data, tree.Root)
	}

	if z != nil {
		w = findParent(z.Data, tree.Root)
	}

	// zig step: If y is the root, do one rotation on x to make it the root
	if z == nil {
		rotate(node, y)
		tree.Root = node
		return
	}

	// zig-zig step: first rotate y and then rotate x
	if (z.RightChild == y && y.RightChild == node) || (z.LeftChild == y && y.LeftChild == node) {
		rotate(y, z)
		rotate(node, y)
	} else { // zig-zag step: double rotate x
		leftSubtree := node.LeftChild
		rightSubtree := node.RightChild
		if node.Data < z.Data {
			node.LeftChild = z.LeftChild
			z.LeftChild = rightSubtree
			node.RightChild = z
			y.RightChild = leftSubtree
		} else {
			node.RightChild = z.RightChild
			z.RightChild = leftSubtree
			node.LeftChild = z
			y.LeftChild = rightSubtree
		}
	}

	if w != nil {
		if w.LeftChild == z {
			w.LeftChild = node
		} else {
			w.RightChild = node
		}
	}

	// recurse if z is not the root of the tree, until node is the root
	if tree.Root != z {
		tree.splay(node)
	}

	tree.Root = node
}

func rotate(x *Node, y *Node) {
	if y.LeftChild == x {
		y.LeftChild = x.RightChild
		x.RightChild = y
	} else {
		y.RightChild = x.LeftChild
		x.LeftChild = y
	}
}

func (tree *SplayTree) PrintTree() {
	printTreeHelper(tree.Root)
}

func printTreeHelper(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Data)
	printTreeHelper(node.LeftChild)
	printTreeHelper(node.RightChild)
}
