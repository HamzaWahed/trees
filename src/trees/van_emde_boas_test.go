package trees

import "testing"

func TestVEB_Insert_baseCase(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(0)

	if tree.Min == NULL || tree.Max == NULL {
		t.Errorf("bug in vEB insertion")
	}
}

func TestBuildVEB_Insert_baseCase2(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(0)
	tree.Insert(1)

	if tree.Min == tree.Max {
		t.Errorf("bug in updating vEB's min and max in thw base case")
		t.Errorf("Min was: %d\nMax was: %d", tree.Min, tree.Max)
	}
}
