package trees

import (
	"fmt"
	"testing"
)

var list1 = []int32{1, 4, 6, 7, 3, 2, 15, 5, 8, 9, 10, 11, 12, 13, 14}

func TestPerfectStaticTree1(t *testing.T) {
	tree := NewStaticTree(list1)
	for _, v := range tree.ToList() {
		fmt.Printf("%d ", v)
	}
}
