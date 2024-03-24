package trees

const (
	RED = iota
	BLACK
)

type RBNode struct {
	color      int
	leftChild  *RBNode
	rightChild *RBNode
	parent     *RBNode
	key        int
}

type RBTree struct {
	root *RBNode
}

func Initialize() *RBTree {
	tree := new(RBTree)
	tree.root = new(RBNode)
	tree.root.color = BLACK
	tree.root.leftChild = nil
	tree.root.rightChild = nil
	tree.root.parent = nil
	return tree
}

func (tree *RBTree) Insert(value int) bool {
	z := new(RBNode)
	if z == nil {
		return false
	}

	z.key = value

	x := tree.root
	var y *RBNode = nil
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.leftChild
		} else {
			x = x.rightChild
		}
	}

	z.parent = y
	if y == nil {
		tree.root = z
	} else if z.key < y.key {
		y.leftChild = z
	} else {
		y.rightChild = z
	}

	z.leftChild = nil
	z.rightChild = nil
	z.color = RED
	tree.insertHelper(z)
	return true
}

func (tree *RBTree) insertHelper(z *RBNode) {
	for z.parent.color == RED {
		if z.parent == z.parent.parent.leftChild {
			y := z.parent.parent.rightChild
			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.rightChild {
					z = z.parent
					tree.leftRotate(z)
				}

				z.parent.color = BLACK
				z.parent.parent.color = RED
				tree.rightRotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.leftChild
			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.leftChild {
					z = z.parent
					tree.rightRotate(z)
				}

				z.parent.color = BLACK
				z.parent.parent.color = RED
				tree.leftRotate(z.parent.parent)
			}
		}
	}

	tree.root.color = BLACK
}

func (tree *RBTree) Delete(z *RBNode) {
	y := z
	var x *RBNode = nil
	yOriginalColor := y.color
	if z.leftChild == nil {
		x = z.rightChild
		tree.transplant(z, z.rightChild)
	} else if z.rightChild == nil {
		x = z.leftChild
		tree.transplant(z, z.leftChild)
	} else {
		y = minimumNode(z.rightChild)
		yOriginalColor = y.color
		x = y.rightChild
		if y != z.rightChild {
			tree.transplant(y, y.rightChild)
			y.rightChild = z.rightChild
			y.rightChild.parent = y
		} else {
			x.parent = y
		}

		tree.transplant(z, y)
		y.leftChild = z.leftChild
		y.leftChild.parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		tree.deleteHelper(x)
	}

}

func (tree *RBTree) deleteHelper(x *RBNode) {
	for x != tree.root && x.color == BLACK {
		if x == x.parent.leftChild {
			w := x.parent.rightChild
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				tree.leftRotate(x.parent)
				w = x.parent.rightChild
			}

			if w.leftChild.color == BLACK && w.rightChild.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.rightChild.color == BLACK {
					w.leftChild.color = BLACK
					w.color = RED
					tree.rightRotate(w)
					w = x.parent.rightChild
				}

				w.color = x.parent.color
				x.parent.color = BLACK
				w.rightChild.color = BLACK
				tree.leftRotate(x.parent)
				x = tree.root
			}
		} else {
			w := x.parent.leftChild
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				tree.rightRotate(x.parent)
				w = x.parent.leftChild
			}

			if w.rightChild.color == BLACK && w.leftChild.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.leftChild.color == BLACK {
					w.rightChild.color = BLACK
					w.color = RED
					tree.leftRotate(w)
					w = x.parent.leftChild
				}

				w.color = x.parent.color
				x.parent.color = BLACK
				w.leftChild.color = BLACK
				tree.rightRotate(x.parent)
				x = tree.root
			}

			x.color = BLACK
		}
	}
}

func (tree *RBTree) transplant(u *RBNode, v *RBNode) {
	if u.parent == nil {
		tree.root = v
	} else if u == u.parent.leftChild {
		u.parent.leftChild = v
	} else {
		u.parent.rightChild = v
	}

	v.parent = u.parent
}

func (tree *RBTree) leftRotate(x *RBNode) {
	y := x.rightChild
	x.rightChild = y.leftChild
	if y.leftChild != nil {
		y.leftChild.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.leftChild {
		x.parent.rightChild = y
	} else {
		x.parent.rightChild = y
	}

	y.leftChild = x
	x.parent = y
}

func (tree *RBTree) rightRotate(x *RBNode) {
	y := x.leftChild
	x.leftChild = y.rightChild
	if y.rightChild != nil {
		y.rightChild.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.rightChild {
		x.parent.rightChild = y
	} else {
		x.parent.leftChild = y
	}

	y.rightChild = x
	x.parent = y
}

func minimumNode(x *RBNode) *RBNode {
	if x.leftChild == nil {
		return x
	}

	return minimumNode(x.leftChild)
}
