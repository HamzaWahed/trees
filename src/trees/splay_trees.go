package trees

import "fmt"

type SplayTree struct {
	Data       int32
	LeftChild  *SplayTree
	RightChild *SplayTree
}

// NewSplayTree Initializes a new splay tree node and returns a pointer to it.
// Runtime: O(1)
func NewSplayTree(x int32) *SplayTree {
	tree := new(SplayTree)
	tree.Data = x
	tree.LeftChild = nil
	tree.RightChild = nil
	return tree
}

// Search Binary searches the splay tree for a node with x in its data field and return true if such a node exists. Splays the
// node to the top if it exists, otherwise splays the node where the search stops on.
// Runtime: O(lg n) amortized
func (tree *SplayTree) Search(x int32) *SplayTree {
	if tree.Data == x {
		return tree
	}

	//TODO: Splay accessed node to the root of the tree
	if tree.LeftChild == nil && tree.RightChild == nil {
		return tree
	}

	if tree.Data < x {
		if tree.RightChild == nil {
			return tree
		}

		return tree.RightChild.Search(x)
	} else {
		if tree.LeftChild == nil {
			return tree
		}

		return tree.LeftChild.Search(x)
	}
}

// Insert Creates a new splay tree node with x and inserts the node into the tree. Splays the node to the root of
// the tree after insertion.
// Runtime:
func (tree *SplayTree) Insert(x int32) bool {
	if tree == nil {
		return false
	}

	parent := tree.Search(x)
	if parent.Data < x {
		parent.RightChild = NewSplayTree(x)
	} else if parent.Data > x {
		parent.LeftChild = NewSplayTree(x)
	}

	//TODO: splay the node to the root of the tree
	//node.splay()
	return true
}

// Not sure if finding parent like this is efficient, or we should be maintaining extra pointers for the data
// structure
func (tree *SplayTree) findParent(x int32) *SplayTree {
	if tree == nil {
		return nil
	}

	if tree.LeftChild == nil && tree.RightChild == nil {
		return nil
	}

	if tree.Data < x {
		if tree.RightChild != nil {
			if tree.RightChild.Data == x {
				return tree
			} else {
				return tree.RightChild.findParent(x)
			}
		}
	} else {
		if tree.LeftChild != nil {
			if tree.LeftChild.Data == x {
				return tree
			} else {
				return tree.LeftChild.findParent(x)
			}
		}
	}

	return nil
}

// maxNode finds the max node in a subtree
// Runtime: O(lg n) worst case
func (tree *SplayTree) maxNode() *SplayTree {
	if tree == nil {
		return nil
	}

	if tree.RightChild == nil {
		return tree
	}

	return tree.RightChild.maxNode()
}

// Delete Deletes the node with x in the splay tree. If the node had a parent, then the parent is splayed to the .
// root of the tree.
// Runtime:
func (tree *SplayTree) Delete(x int32) (int32, string) {
	nodeToDelete := tree.Search(x)
	err := ""
	if nodeToDelete == nil || nodeToDelete.Data != x {
		err = fmt.Sprintf("{x} is not in the splay tree")
		return x, err
	}

	var leftSubtree, rightSubtree *SplayTree = nil, nil

	// deleting the node x is in may create two subtrees of the children
	// add pointers to these nodes, so they can be merged with the main tree afterward.
	if nodeToDelete.LeftChild != nil {
		leftSubtree = nodeToDelete.LeftChild
	}

	if nodeToDelete.RightChild != nil {
		rightSubtree = nodeToDelete.RightChild
	}

	// disconnect the main tree from the node x is in
	parent := tree.findParent(x)
	if parent.LeftChild.Data == x {
		parent.LeftChild = nil
	} else {
		parent.RightChild = nil
	}

	//TODO: find the max element in the left subtree and splay it to the top of that tree
	maxNodeLeftSubtree := leftSubtree.maxNode()
	maxNodeLeftSubtree.splay()

	// connect the left and right subtrees, and then connect them to the main tree
	leftSubtree.merge(rightSubtree)
	parent.merge(leftSubtree)

	// splay the parent node to the top
	parent.splay()

	return x, err
}

func (tree *SplayTree) merge(treeToMerge *SplayTree) {

}

func (tree *SplayTree) splay() {

}

func (tree *SplayTree) PrintTree() {
	if tree == nil {
		return
	}

	fmt.Println(tree.Data)
	if tree.LeftChild != nil {
		tree.LeftChild.PrintTree()
	}

	if tree.RightChild != nil {
		tree.RightChild.PrintTree()
	}
}
