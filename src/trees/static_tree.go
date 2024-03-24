package trees

import (
	"slices"
)

type StaticTree struct {
	root *Node
	size int32
}

func NewStaticTree(keys []int32) *StaticTree {
	tree := new(StaticTree)

	if len(keys) < 1 {
		panic("Key list is empty")
	}
	slices.SortFunc(keys, func(a, b int32) int {
		return int(a - b)
	})

	InsertKeyList(tree, keys)
	return tree
}

func InsertKeyList(tree *StaticTree, keys []int32) {
	if tree == nil {
		panic("Cant insert into a null tree")
	} else if len(keys) == 0 {
		return
	}

	var key_to_insert = 1
	for key_to_insert <= len(keys)/2 {
		key_to_insert *= 2
	}

	//not sure how this is diff from just having the for loop end at <, but it works while the other doesnt
	if key_to_insert/2-1 <= (len(keys) - key_to_insert) {
		key_to_insert--
	} else {
		key_to_insert = len(keys) - key_to_insert/2
	}

	node := new(Node)
	node.Data = keys[key_to_insert]

	//edge case where the tree is initially empty
	if tree.root == nil {
		tree.root = node
	} else {
		InsertNode(tree.root, node)
	}

	// recursively calling the function on the rest of the key list
	InsertKeyList(tree, keys[:key_to_insert])
	InsertKeyList(tree, keys[key_to_insert+1:])

}

// InsertNode : normal BST insertion. returns true if the node was added, and false if it's already in the tree/**
func InsertNode(root *Node, node *Node) bool {
	if node == nil {
		panic("Cant insert null node")
	}

	if root == nil {
		root = node
		return true
	}

	if root.Data < node.Data {
		if root.RightChild == nil {
			root.RightChild = node
			return true
		} else {
			InsertNode(root.RightChild, node)
		}
	} else if root.Data > node.Data {
		if root.LeftChild == nil {
			root.LeftChild = node
			return true
		} else {
			InsertNode(root.LeftChild, node)
		}
	}

	return false
}

func (tree *StaticTree) ToList() []int32 {
	if tree.root == nil {
		return nil
	}

	var values []int32
	var node *Node
	queue := make([]*Node, 0)
	queue = append(queue, tree.root)

	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		values = append(values, node.Data)
		if node.LeftChild != nil {
			queue = append(queue, node.LeftChild)
		}

		if node.RightChild != nil {
			queue = append(queue, node.RightChild)
		}
	}
	return values
}
