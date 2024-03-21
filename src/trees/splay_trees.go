package trees

const (
	rightChild = iota
	leftChild
	parent
	disconnected
)

type Node struct {
	Data       int32
	LeftChild  *Node
	RightChild *Node
	parent     *Node
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
	tree.Root.parent = nil
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
			node.parent = root
		} else {
			insertHelper(root.LeftChild, node)
		}
	} else {
		if root.RightChild == nil {
			root.RightChild = node
			node.parent = root
		} else {
			insertHelper(root.RightChild, node)
		}
	}
}

// Returns the relation between y and x
func isChild(x *Node, y *Node) int {
	if x == nil || y == nil {
		panic("isChild Operation on Nil nodes")
	}

	if y.LeftChild == x {
		return leftChild
	} else if y.RightChild == x {
		return rightChild
	} else {
		return disconnected
	}
}

func (tree *SplayTree) splay(node *Node) {

	//checks if node is the root
	y := node.parent
	if y == nil {
		return
	}

	var z *Node = nil
	if y != nil {
		z = y.parent
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

	// recurse if z is not the root of the tree, until node is the root
	if tree.Root != z {
		tree.splay(node)
	}

	tree.Root = node
}

// rotate single rotation on x and disperses subtrees of x's children to y
func rotate(x *Node, y *Node) {

	if y.LeftChild == x {
		if x.RightChild != nil {
			x.RightChild.parent = y
		}
		y.LeftChild = x.RightChild
		x.RightChild = y
	} else {
		if x.LeftChild != nil {
			x.LeftChild.parent = y
		}
		y.RightChild = x.LeftChild
		x.LeftChild = y
	}

	if y.parent != nil {
		var yPosition int = isChild(y, y.parent)

		if yPosition == leftChild {
			y.parent.LeftChild = x
		} else if yPosition == rightChild {
			y.parent.RightChild = x
		}
	}

	x.parent = y.parent
	y.parent = x
}

// doubleRotate used during the zig-zag step to do a double rotation on x and disperses the subtrees of x's children
// to y and z
func doubleRotate(x *Node, y *Node, z *Node) {
	rotate(x, y)
	rotate(x, z)
}

func (tree *SplayTree) Delete(val int32) {
	x := findNode(tree.Root, val)
	parent := x.parent

	//deleting all references to x, and dividing the tree into 3
	rightSubtree := new(SplayTree)
	rightSubtree.Root = x.RightChild
	x.RightChild = nil
	if rightSubtree.Root != nil {
		rightSubtree.Root.parent = nil
	}

	leftSubtree := new(SplayTree)
	leftSubtree.Root = x.LeftChild
	x.LeftChild = nil
	if leftSubtree.Root != nil {
		leftSubtree.Root.parent = nil
		leftSubtree.splay(leftSubtree.findMaxNode())
		leftSubtree.Root.RightChild = rightSubtree.Root
	} else {
		leftSubtree = rightSubtree
	}

	if parent != nil {
		if parent.LeftChild == x {
			parent.LeftChild = nil
		} else {
			parent.RightChild = nil
		}
		x.parent = nil
		tree.splay(parent)

		if parent.Data < leftSubtree.Root.Data {
			parent.RightChild = leftSubtree.Root
		} else {
			parent.LeftChild = leftSubtree.Root
		}
		tree.Root = parent
	} else {
		tree.Root = leftSubtree.Root
	}

	tree.Size--
}

// finds the node holding the element to delete, and nil if the element is not in the tree
func findNode(root *Node, val int32) *Node {
	if root == nil {
		return nil
	}
	if root.Data == val {
		return root
	} else if root.Data < val {
		return findNode(root.RightChild, val)
	} else {
		return findNode(root.LeftChild, val)
	}
}

// finds the largest element in a subtree
func (tree *SplayTree) findMaxNode() *Node {
	var root = tree.Root
	if root == nil {
		panic("Calling max node on an empty tree")
	}

	for root.RightChild != nil {
		root = root.RightChild
	}
	return root
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
