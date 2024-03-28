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
