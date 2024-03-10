package trees

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

	tree.Root.Data = x
	tree.Size = 1
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
	tree.Size++
	tree.splay(node)
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
	if y == nil {
		return
	}

	var z *Node = nil
	var w *Node = nil
	if y != nil {
		z = findParent(y.Data, tree.Root)
	}

	// may have to update z parent's pointer
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
		doubleRotate(node, y, z)
	}

	// update z parent's pointer after zig-zig or zig-zag step if the parent exists
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

// rotate single rotation on x and disperses subtrees of x's children to y
func rotate(x *Node, y *Node) {
	if y.LeftChild == x {
		y.LeftChild = x.RightChild
		x.RightChild = y
	} else {
		y.RightChild = x.LeftChild
		x.LeftChild = y
	}
}

// doubleRotate used during the zig-zag step to do a double rotation on x and disperses the subtrees of x's children
// to y and z
func doubleRotate(x *Node, y *Node, z *Node) {
	leftSubtree := x.LeftChild
	rightSubtree := x.RightChild
	if x.Data < z.Data {
		x.LeftChild = z.LeftChild
		z.LeftChild = rightSubtree
		x.RightChild = z
		y.RightChild = leftSubtree
	} else {
		x.RightChild = z.RightChild
		z.RightChild = leftSubtree
		x.LeftChild = z
		y.LeftChild = rightSubtree
	}
}

// ToList converts the splay tree to a splice
func (tree *SplayTree) ToList() []int32 {
	if tree.Root == nil {
		return nil
	}

	var values []int32
	var node *Node
	queue := make([]*Node, 0)
	queue = append(queue, tree.Root)

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
