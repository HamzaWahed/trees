package trees

import (
	"fmt"
	"slices"
)

const (
	RED = iota
	BLACK
)

// TangoTree are O(lg lg n) competitive Binary Search Trees that has a trees-of-trees representation. These
// smaller trees of height O(lg lg n) are called Auxiliary Trees. Auxiliary trees are augmented Red-Black trees
// that supports link and cut operations for updating the tango trees.
type TangoTree struct {
	Root *AuxNode
}

type AuxNode struct {
	Color             int
	LeftChild         *AuxNode
	RightChild        *AuxNode
	Parent            *AuxNode
	PreferredChild    *AuxNode
	Key               int
	DepthInP          int
	IsRoot            bool
	MaxDepthInSubtree int
	MinDepthInSubtree int
}

// NewTangoTree creates a new Tango Tree with an array of integers.
func NewTangoTree(keys []int) *TangoTree {
	tree := new(TangoTree)

	if len(keys) < 1 {
		panic("Key list is empty")
	}
	slices.SortFunc(keys, func(a, b int) int {
		return int(a - b)
	})

	insertKeyList(tree, keys)
	return tree
}

func insertKeyList(tree *TangoTree, keys []int) {
	if tree == nil {
		panic("Cant insert into a null tree")
	} else if len(keys) == 0 {
		return
	}

	var key_to_insert = 1
	for key_to_insert <= len(keys)/2 {
		key_to_insert *= 2
	}

	//not sure how this is diff from just having the for loop end at <, but it works while the other doesn't
	if key_to_insert/2-1 <= (len(keys) - key_to_insert) {
		key_to_insert--
	} else {
		key_to_insert = len(keys) - key_to_insert/2
	}

	node := new(AuxNode)
	node.IsRoot = true
	node.Color = BLACK
	node.Key = keys[key_to_insert]

	//edge case where the tree is initially empty
	if tree.Root == nil {
		node.DepthInP = 0
		node.MinDepthInSubtree = 0
		node.MaxDepthInSubtree = 0
		node.Parent = nil
		tree.Root = node
	} else {
		node.DepthInP = tree.Root.DepthInP + 1
		insertNode(tree.Root, node)
	}

	// recursively calling the function on the rest of the key list
	insertKeyList(tree, keys[:key_to_insert])
	insertKeyList(tree, keys[key_to_insert+1:])

}

// insertNode : normal BST insertion. returns true if the node was added, and false if it's already in the tree/**
func insertNode(root *AuxNode, node *AuxNode) bool {
	if node == nil {
		panic("Cant insert null node")
	}

	// set the depth of node being inserted
	node.DepthInP = root.DepthInP + 1
	node.MinDepthInSubtree = node.DepthInP
	node.MaxDepthInSubtree = node.DepthInP

	if root.Key < node.Key {
		if root.RightChild == nil {
			root.RightChild = node
			return true
		} else {
			insertNode(root.RightChild, node)
		}
	} else if root.Key > node.Key {
		if root.LeftChild == nil {
			root.LeftChild = node
			return true
		} else {
			insertNode(root.LeftChild, node)
		}
	}

	return false
}

func PrintAuxNodes(node *AuxNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Color)
	PrintAuxNodes(node.LeftChild)
	PrintAuxNodes(node.RightChild)
}

// Access does a binary search while updating preferred child pointers.
func (tree *TangoTree) Access(node *AuxNode, key int) *AuxNode {
	if node.Key == key {
		return node
	}

	if node.Key > key {
		if node.PreferredChild == nil {
			node.PreferredChild = node.LeftChild
		} else if node.PreferredChild != node.LeftChild {
			tango(tree.Root, node)
		}

		return tree.Access(node.LeftChild, key)
	}

	if node.PreferredChild == nil {
		node.PreferredChild = node.RightChild
	} else if node.PreferredChild != node.RightChild {
		tango(tree.Root, node)
	}

	return tree.Access(node.RightChild, key)
}

func tango(rootAuxTree *AuxNode, auxTreeToMerge *AuxNode) {

}

//func (tree *AuxTree) Insert(ttNode *TangoTreeNode) bool {
//	z := new(AuxNode)
//	z.Key = ttNode.Key
//
//	if z == nil {
//		return false // failed to allocate a new AuxNode
//	}
//
//	z.value.Key = tt_node.Key
//
//	x := tree.Root
//	var y *RBNode = nil
//	for x != nil {
//		y = x
//		if z.value.Key < x.value.Key {
//			x = x.leftChild
//		} else {
//			x = x.rightChild
//		}
//	}
//
//	z.parent = y
//	if y == nil {
//		tree.Root = z
//	} else if z.value.Key < y.value.Key {
//		y.leftChild = z
//	} else {
//		y.rightChild = z
//	}
//
//	z.leftChild = nil
//	z.rightChild = nil
//	z.color = RED
//	tree.insertHelper(z)
//	return true
//}
//
//func (tree *RBTree) insertHelper(z *RBNode) {
//	for z.parent.color == RED {
//		if z.parent == z.parent.parent.leftChild {
//			y := z.parent.parent.rightChild
//			if y.color == RED {
//				z.parent.color = BLACK
//				y.color = BLACK
//				z.parent.parent.color = RED
//				z = z.parent.parent
//			} else {
//				if z == z.parent.rightChild {
//					z = z.parent
//					tree.leftRotate(z)
//				}
//
//				z.parent.color = BLACK
//				z.parent.parent.color = RED
//				tree.rightRotate(z.parent.parent)
//			}
//		} else {
//			y := z.parent.parent.leftChild
//			if y.color == RED {
//				z.parent.color = BLACK
//				y.color = BLACK
//				z.parent.parent.color = RED
//				z = z.parent.parent
//			} else {
//				if z == z.parent.leftChild {
//					z = z.parent
//					tree.rightRotate(z)
//				}
//
//				z.parent.color = BLACK
//				z.parent.parent.color = RED
//				tree.leftRotate(z.parent.parent)
//			}
//		}
//	}
//
//	tree.Root.color = BLACK
//}
//
//func (tree *RBTree) Delete(z *RBNode) {
//	y := z
//	var x *RBNode = nil
//	yOriginalColor := y.color
//	if z.leftChild == nil {
//		x = z.rightChild
//		tree.transplant(z, z.rightChild)
//	} else if z.rightChild == nil {
//		x = z.leftChild
//		tree.transplant(z, z.leftChild)
//	} else {
//		y = minimumNode(z.rightChild)
//		yOriginalColor = y.color
//		x = y.rightChild
//		if y != z.rightChild {
//			tree.transplant(y, y.rightChild)
//			y.rightChild = z.rightChild
//			y.rightChild.parent = y
//		} else {
//			x.parent = y
//		}
//
//		tree.transplant(z, y)
//		y.leftChild = z.leftChild
//		y.leftChild.parent = y
//		y.color = z.color
//	}
//
//	if yOriginalColor == BLACK {
//		tree.deleteHelper(x)
//	}
//
//}
//
//func (tree *RBTree) deleteHelper(x *RBNode) {
//	for x != tree.Root && x.color == BLACK {
//		if x == x.parent.leftChild {
//			w := x.parent.rightChild
//			if w.color == RED {
//				w.color = BLACK
//				x.parent.color = RED
//				tree.leftRotate(x.parent)
//				w = x.parent.rightChild
//			}
//
//			if w.leftChild.color == BLACK && w.rightChild.color == BLACK {
//				w.color = RED
//				x = x.parent
//			} else {
//				if w.rightChild.color == BLACK {
//					w.leftChild.color = BLACK
//					w.color = RED
//					tree.rightRotate(w)
//					w = x.parent.rightChild
//				}
//
//				w.color = x.parent.color
//				x.parent.color = BLACK
//				w.rightChild.color = BLACK
//				tree.leftRotate(x.parent)
//				x = tree.Root
//			}
//		} else {
//			w := x.parent.leftChild
//			if w.color == RED {
//				w.color = BLACK
//				x.parent.color = RED
//				tree.rightRotate(x.parent)
//				w = x.parent.leftChild
//			}
//
//			if w.rightChild.color == BLACK && w.leftChild.color == BLACK {
//				w.color = RED
//				x = x.parent
//			} else {
//				if w.leftChild.color == BLACK {
//					w.rightChild.color = BLACK
//					w.color = RED
//					tree.leftRotate(w)
//					w = x.parent.leftChild
//				}
//
//				w.color = x.parent.color
//				x.parent.color = BLACK
//				w.leftChild.color = BLACK
//				tree.rightRotate(x.parent)
//				x = tree.Root
//			}
//
//			x.color = BLACK
//		}
//	}
//}
//
//func (tree *RBTree) transplant(u *RBNode, v *RBNode) {
//	if u.parent == nil {
//		tree.Root = v
//	} else if u == u.parent.leftChild {
//		u.parent.leftChild = v
//	} else {
//		u.parent.rightChild = v
//	}
//
//	v.parent = u.parent
//}
//
//func (tree *RBTree) leftRotate(x *RBNode) {
//	y := x.rightChild
//	x.rightChild = y.leftChild
//	if y.leftChild != nil {
//		y.leftChild.parent = x
//	}
//
//	y.parent = x.parent
//	if x.parent == nil {
//		tree.Root = y
//	} else if x == x.parent.leftChild {
//		x.parent.rightChild = y
//	} else {
//		x.parent.rightChild = y
//	}
//
//	y.leftChild = x
//	x.parent = y
//}
//
//func (tree *RBTree) rightRotate(x *RBNode) {
//	y := x.leftChild
//	x.leftChild = y.rightChild
//	if y.rightChild != nil {
//		y.rightChild.parent = x
//	}
//
//	y.parent = x.parent
//	if x.parent == nil {
//		tree.Root = y
//	} else if x == x.parent.rightChild {
//		x.parent.rightChild = y
//	} else {
//		x.parent.leftChild = y
//	}
//
//	y.rightChild = x
//	x.parent = y
//}
//
//func minimumNode(x *RBNode) *RBNode {
//	if x.leftChild == nil {
//		return x
//	}
//
//	return minimumNode(x.leftChild)
//}
