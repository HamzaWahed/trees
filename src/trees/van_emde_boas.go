package trees

import "math"

type VEB struct {
	Summary          *VEB
	UniverseSize     int // assumed to be a power of 2
	LowerSquareRoot  int
	HigherSquareRoot int
	Cluster          []*VEB
	Min              int
	Max              int
}

const NULL = -1

// BuildVEB recursively builds the Van Emde Boas tree of the specified universe size. Current implementation only
// works for universe size of powers of 2.
func BuildVEB(universeSize int) *VEB {
	tree := new(VEB)
	tree.UniverseSize = universeSize
	tree.LowerSquareRoot = int(math.Pow(2, math.Floor(math.Log2(float64(universeSize))/2)))
	tree.HigherSquareRoot = int(math.Pow(2, math.Ceil(math.Log2(float64(universeSize))/2)))
	tree.Max = NULL
	tree.Min = NULL

	if universeSize == 2 {
		return tree
	}

	tree.Summary = BuildVEB(tree.HigherSquareRoot)
	tree.Cluster = make([]*VEB, tree.HigherSquareRoot)
	for i := range tree.Cluster {
		tree.Cluster[i] = BuildVEB(tree.LowerSquareRoot)
	}

	return tree
}

// Maximum returns the max field of the Van Emde Boas Tree structure
func (tree *VEB) Maximum() int {
	return tree.Max
}

// Minimum returns the min field of the Van Emde Boas Tree structure
func (tree *VEB) Minimum() int {
	return tree.Min
}

// Member searches if x is a member of the Van Emde Boas Tree.
func (tree *VEB) Member(x int) bool {
	if x == tree.Min || x == tree.Max {
		return true
	}

	if tree.UniverseSize == 2 {
		return false
	}

	cluster := tree.Cluster[tree.high(x)]
	return cluster.Member(tree.low(x))
}

// Successor Searches for the successor of x in the Van Emde Boas Tree. Returns -1 if the successor does not exist
// in the universe of the Root Van Emde Boas Tree.
func (tree *VEB) Successor(x int) int {
	if tree.UniverseSize == 2 {
		if x == 0 && tree.UniverseSize == 2 {
			return 1
		}

		return NULL
	} else if tree.Min != NULL && x < tree.Min {
		return tree.Min
	}

	cluster := tree.Cluster[tree.high(x)]
	maxLow := cluster.Max
	lowerBits := tree.low(x)
	if maxLow != NULL && lowerBits < maxLow {
		offset := cluster.Successor(lowerBits)
		return tree.index(tree.high(x), offset)
	}

	succCluster := tree.Summary.Successor(tree.high(x))
	if succCluster == NULL {
		return NULL
	}

	offset := tree.Cluster[succCluster].Min
	return tree.index(succCluster, offset)
}

// Predecessor Searches for the predecessor of x in the Van Emde Boas Tree. Returns -1 if the predecessor does not exist
// in the universe of the Root Van Emde Boas Tree.
func (tree *VEB) Predecessor(x int) int {
	if tree.UniverseSize == 2 {
		if x == 1 && tree.Min == 0 {
			return 0
		}

		return NULL
	}

	if tree.Max != NULL && x > tree.Max {
		return tree.Max
	}

	cluster := tree.Cluster[tree.high(x)]
	minLow := cluster.Min
	lowerBits := tree.low(x)
	if minLow != NULL && lowerBits > minLow {
		offset := cluster.Predecessor(lowerBits)
		return tree.index(tree.high(x), offset)
	}

	predCluster := tree.Summary.Predecessor(tree.high(x))
	if predCluster == NULL {
		if tree.Min != NULL && x > tree.Min {
			return tree.Min
		}

		return NULL
	}

	offset := tree.Cluster[predCluster].Max
	return tree.index(predCluster, offset)
}

// emptyTreeInsert updates the min and max fields of the Van Emde Boas Tree.
func (tree *VEB) emptyTreeInsert(x int) {
	tree.Min = x
	tree.Max = x
}

// Insert inserts an element in the Van Emde Boas Tree. Swaps the element to be inserted with the min field and inserts
// the previous minimum element if the element to be inserted is less than the current minimum.
func (tree *VEB) Insert(x int) {
	if tree.Min == NULL {
		tree.emptyTreeInsert(x)
	} else if x < tree.Min {
		x, tree.Min = tree.Min, x
		if tree.UniverseSize > 2 {
			clusterIndex := tree.high(x)
			if tree.Cluster[clusterIndex].Min == NULL {
				tree.Summary.Insert(clusterIndex)
				tree.Cluster[clusterIndex].emptyTreeInsert(tree.low(x))
			} else {
				tree.Cluster[tree.high(x)].Insert(tree.low(x))
			}
		}

		if x > tree.Max {
			tree.Max = x
		}
	}
}

func (tree *VEB) delete(x int) {
	//if tree.Member(x) == false {
	//	return
	//}

	if tree.Min == tree.Max {
		tree.Min = NULL
		tree.Max = NULL
	} else if tree.UniverseSize == 2 {
		if x == 0 {
			tree.Min = 1
		} else {
			tree.Min = 0
		}

		tree.Max = tree.Min
	} else {

		//handling the special case where x is the min, here we replace x with the second-smallest element in the
		// cluster, set Min to the new element then change Min to the new element's value (thus deleting x in the process)
		if x == tree.Min {
			var firstCluster int = tree.Summary.Minimum()
			var offset = tree.Cluster[firstCluster].Min
			x = tree.index(firstCluster, offset)
			tree.Min = x
		}

		//deleting x from its cluster
		tree.Cluster[tree.high(x)].delete(tree.low(x))

		//checking if the cluster is now empty, here we check if min == NULL because if so then max would also be NUll and the cluster is empty
		if tree.Cluster[tree.high(x)].Min == NULL {
			tree.Summary.delete(tree.high(x))
		}

		//since we could've changed the value of x as mention above, we now need to check if x is the tree's max
		if x == tree.Max {
			var maxSummary int = tree.Summary.Max

			// if tree only has a min
			if maxSummary == NULL {
				tree.Max = tree.Min
			} else {
				var offset int = tree.Cluster[maxSummary].Max
				tree.Max = tree.index(maxSummary, offset)
			}
		}

		if x == tree.Max {
			var offset int = tree.Cluster[tree.high(x)].Max
			tree.Max = tree.index(tree.high(x), offset)
		}
	}
}

// finds the index of the element in the Van Emde Boas Tree with respect to the universe size of
// the Van Emde Boas Tree passed in
func (tree *VEB) index(x int, y int) int {
	return x*tree.LowerSquareRoot + y
}

// finds the high bits of x
func (tree *VEB) high(x int) int {
	return int(math.Floor(float64(x) / float64(tree.LowerSquareRoot)))
}

// finds the low bits of x
func (tree *VEB) low(x int) int {
	return x % tree.LowerSquareRoot
}
