package trees

type MultiSplayNode struct {
	leftChild  *MultiSplayNode
	rightChild *MultiSplayNode
	parent     *MultiSplayNode
	refDepth   int
	minDepth   int
	isRoot     bool
}
