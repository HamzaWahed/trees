package trees

import "slices"

type TangoTreeNode struct {
	LeftChild         *TangoTreeNode
	RightChild        *TangoTreeNode
	Parent            *TangoTreeNode
	PreferredChild    *TangoTreeNode
	Key               int32
	IsRoot            bool
	DepthInStaticTree int32
}

type TangoTree struct {
	root *TangoTreeNode
}

// BuildTangoTree builds a Tango tree with the array of integers passed in.
func BuildTangoTree(keys []int32) *TangoTree {
	tree := new(TangoTree)

	if len(keys) < 1 {
		panic("Key list is empty")
	}
	slices.SortFunc(keys, func(a, b int32) int {
		return int(a - b)
	})

	insertKeyList(tree, keys)
	return tree
}

// insertKeyList inserts a list of keys into the tango tree when building it for the first time.
func insertKeyList(tree *TangoTree, keys []int32) {
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

	node := new(TangoTreeNode)
	node.Key = keys[key_to_insert]

	//edge case where the tree is initially empty
	if tree.root == nil {
		node.DepthInStaticTree = 1
		tree.root = node
	} else {
		insertNode(tree.root, node, tree.root.DepthInStaticTree+1)
	}

	// recursively calling the function on the rest of the key list
	insertKeyList(tree, keys[:key_to_insert])
	insertKeyList(tree, keys[key_to_insert+1:])

}

// InsertNode : normal BST insertion. returns true if the node was added, and false if it's already in the tree/**
func insertNode(root *TangoTreeNode, node *TangoTreeNode, depth int32) bool {
	if node == nil {
		panic("Cant insert null node")
	}
	node.DepthInStaticTree = depth

	if root.Key < node.Key {
		if root.RightChild == nil {
			root.RightChild = node
			node.Parent = root
			return true
		} else {
			insertNode(root.RightChild, node, depth+1)
		}
	} else if root.Key > node.Key {
		if root.LeftChild == nil {
			root.LeftChild = node
			node.Parent = root
			return true
		} else {
			insertNode(root.LeftChild, node, depth+1)
		}
	}

	return false
}

// Search searches for a key in the Tango tree. As it walks down the tree, it updates the preferred child pointers
// to match the path it is following.
func (tree *TangoTree) Search(key int32) *TangoTreeNode {
	if tree.root == nil {
		return nil
	}

	tree.root.IsRoot = true
	node := tt_searchHelper(tree.root, key)
	updateTreeStructure(tree, tree.root)
	return node
}

func tt_searchHelper(node *TangoTreeNode, key int32) *TangoTreeNode {
	if node.Key == key {
		return node
	}

	if node.Key < key && node.LeftChild != nil {
		node.PreferredChild = node.LeftChild
		return tt_searchHelper(node.LeftChild, key)
	} else if node.Key > key && node.RightChild != nil {
		node.PreferredChild = node.RightChild
		return tt_searchHelper(node.RightChild, key)
	}

	return nil
}

func updateTreeStructure(tree *TangoTree, node *TangoTreeNode) {
	// find preferred paths and create aux tree
}
