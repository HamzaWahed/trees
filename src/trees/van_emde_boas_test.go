package trees

import (
	"math"
	"testing"
)

func TestIndexing_high_low(t *testing.T) {
	var u int = 128
	var x int = 33
	//var lower_sqrt = 8
	//var upper_sqrt = 16
	//var true_high int = 11
	//var true_low int = 0

	var h int = high(x, u)
	var l int = low(x, u)

	if h != 4 || l != 1 {
		t.Errorf("Incorrect Indexing:\nHigh should be 4, you got %d\nLow should be 1 you got %d", h, l)
	}
}

func TestIndexing_index(t *testing.T) {
	var u int = 128
	//var x int = 33
	var h int = 4
	var lower_sqrt int = 8
	var l int = 1

	if index(h, l, u) != h*lower_sqrt+l {
		t.Errorf("The correct index should be 13 we got %d", int(math.Pow(2, math.Floor(math.Log(float64(u))/2))))
	}

}
