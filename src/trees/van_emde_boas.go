package trees

import "math"

type vEB struct {
	Summary          *vEB
	U                int
	lowerSquareRoot  int
	higherSquareRoot int
	Cluster          []*vEB
	Min              int
	Max              int
}

const NULL = -1

func low(x int, u int) int {
	lfloor := int(math.Pow(2, math.Floor(math.Log(float64(u)/2))))
	return x % lfloor
}

func high(x int, u int) int {
	lfloor := int(math.Pow(2, math.Floor(math.Log(float64(u)/2))))
	return int(math.Floor(float64(x / lfloor)))
}

func index(x int, y int, u int) int {
	lfloor := int(math.Pow(2, math.Floor(math.Log(float64(u)/2))))
	return x*lfloor + y
}

func Maximum(V *vEB) int {
	return V.Max
}

func Minimum(V *vEB) int {
	return V.Min
}

func Member(V *vEB, x int) bool {
	if x == V.Min || x == V.Max {
		return true
	}

	if V.U == 2 {
		return false
	}

	return Member(V.Cluster[high(x, V.U)], low(x, V.U))
}

func Successor(V *vEB, x int) int {
	if V.U == 2 {
		if x == 0 && V.Max == 1 {
			return 1
		}

		return NULL
	} else if V.Min != NULL && x < V.Min {
		return V.Min
	}

	maxLow := Maximum(V.Cluster[high(x, V.U)])
	if maxLow != NULL && low(x, V.U) < maxLow {
		offset := Successor(V.Cluster[high(x, V.U)], low(x, V.U))
		return index(high(x, V.U), offset, V.Cluster[high(x, V.U)].U)
	}

	succCluster := Successor(V.Summary, high(x, V.U))
	if succCluster == NULL {
		return NULL
	}

	offset := Minimum(V.Cluster[succCluster])
	return index(succCluster, offset, V.Cluster[succCluster].U)
}

func Predecessor(V *vEB, x int) int {
	if V.U == 2 {
		if x == 1 && V.Min == 0 {
			return 0
		}

		return NULL
	}

	if V.Max != NULL && x > V.Max {
		return V.Max
	}

	minLow := Minimum(V.Cluster[high(x, V.U)])
	if minLow != NULL && low(x, V.U) > minLow {
		offset := Predecessor(V.Cluster[high(x, V.U)], low(x, V.U))
		return index(high(x, V.U), offset, V.Cluster[high(x, V.U)].U)
	}

	predCluster := Predecessor(V.Summary, high(x, high(x, V.U)))
	if predCluster == NULL {
		if V.Min != NULL && x > V.Min {
			return V.Min
		}

		return NULL
	}

	offset := Maximum(V.Cluster[predCluster])
	return index(predCluster, offset, V.Cluster[predCluster].U)

}
