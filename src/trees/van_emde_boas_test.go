package trees

import "testing"

func vEBUniverse16() *VEB {
	var tree = BuildVEB(16)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(1)
	return tree
}

func vEBUniverse512() *VEB {
	var tree = BuildVEB(1024)
	tree.Insert(14)
	tree.Insert(11)
	tree.Insert(105)
	tree.Insert(15)
	tree.Insert(214)
	tree.Insert(6)
	tree.Insert(7)
	return tree
}

func BenchmarkTestVEB_Insert_baseCase(t *testing.B) {
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

func TestBuildVEB_Insert_baseCase3(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(1)
	tree.Insert(0)

	if tree.Min == tree.Max {
		t.Errorf("bug when inserting max followed by min\nMin is %d\nMax is %d", tree.Min, tree.Max)
	}
}

func TestBuildVEB_Insert_duplicateBaseCase(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(0)
	tree.Insert(0)

	if tree.Min != tree.Max || tree.Min == NULL || tree.Max == NULL {
		t.Errorf("Error when inserting the same element to the same empty tree")
	}
}

func TestBuildVEB_Delete_fromBaseCase_with_1_element(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(0)
	tree.delete(0)

	if tree.Min != tree.Max && tree.Min != NULL {
		t.Errorf("Error in deleting an element from a base tree with 1 element")
	}
}

func TestBuildVEB_Delete_fromBaseCase_with_2_elements(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(0)
	tree.Insert(1)
	tree.delete(1)

	if tree.Min != tree.Max || tree.Max != 0 {
		t.Errorf("Error when deleting an elemnet from a base case tree with 2 eleemnts in it")
	}
}

func TestBuildVEB_Delete_fromBaseCase_inEmptyTree(t *testing.T) {
	var tree = BuildVEB(2)
	tree.delete(0)

	if tree.Min != tree.Max && tree.Min != NULL {
		t.Errorf("Deleting from an empty base case tree should result in both min and max being NULL(-1)")
	}
}

func TestBuildVEB_Predecessor_baseCase1(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(0)
	tree.Insert(1)

	if tree.Predecessor(0) != NULL && tree.Predecessor(1) != 0 {
		t.Errorf("Error when search for the predecessor in the base case")
	}
}

func TestBuildVEB_Predecessor_baseCase2(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(0)

	if tree.Predecessor(1) != 0 {
		t.Errorf("Error when searching for the predecessor of an element that isnt in the tree")
	}
}

func TestBuildVEB_Successor_baseCase1(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(0)
	tree.Insert(1)

	if tree.Successor(0) != 1 && tree.Predecessor(1) != NULL {
		t.Errorf("Error when search for the successor in the base case")
	}
}

func TestBuildVEB_Successor_baseCase2(t *testing.T) {
	var tree = BuildVEB(2)
	tree.Insert(1)

	if tree.Successor(0) != 1 {
		t.Errorf("Error when searching for the successor of an element that isnt in the tree")
	}
}

func TestBuildVEB_recursive_Insert1(t *testing.T) {
	var tree = BuildVEB(4)
	tree.Insert(0)
	tree.Insert(2)

	if tree.Min != 0 && tree.Max != 2 {
		t.Errorf("Error when inserting 2 elements in different clusters")
	}
	if tree.Cluster[0].Min != tree.Cluster[0].Max && tree.Cluster[0].Min != 0 {
		t.Errorf("Error in recurize insertion of lowest element")
	}
	if tree.Cluster[1].Min != tree.Cluster[1].Max && tree.Cluster[1].Min != 0 {
		t.Errorf("Error in recurize insertion of highest element")
	}
}

func TestBuildVEB_recursive_Insert2(t *testing.T) {
	tree := vEBUniverse16()

	if tree.Min != 1 && tree.Max != 12 {
		t.Errorf("Error in updating min and max of a recursive vEM tree")
	}

	a := tree.Min
	b := tree.Cluster[0].Min
	//c's relative position is 3
	c := tree.Cluster[1].Max
	d := tree.Max
	if a != 1 || b != 3 || c != 3 || d != 12 {
		t.Errorf("Inserted elements are not present in the recurisve vEB structure\na=%d\nb=%d\nc=%d\nd=%d", a, b, c, d)
	}
}

func TestBuildVEB_recursive_Membership1(t *testing.T) {
	tree := vEBUniverse16()

	if !(tree.Member(3) && tree.Member(7) && tree.Member(1) && tree.Member(12)) {
		t.Errorf("Error in recursive membership")
	}
}

func TestBuildVEB_recursive_Membership2(t *testing.T) {
	tree := vEBUniverse512()

	if !(tree.Member(15) && tree.Member(105) && tree.Member(11) && tree.Member(214)) {
		t.Errorf("Error in recursive membership")
	}
}

func TestBuildVEB_recursive_Predecessor(t *testing.T) {
	tree := vEBUniverse512()

	if tree.Predecessor(214) != 105 && tree.Predecessor(105) != 15 && tree.Predecessor(6) != NULL {
		t.Errorf("Error in predecessor search for recursive vEBs")
	}
}

func TestBuildVEB_recursive_Successor(t *testing.T) {
	tree := vEBUniverse512()

	if tree.Successor(7) != 11 && tree.Successor(15) != 105 && tree.Successor(214) != NULL {
		t.Errorf("Error in successor search for recursive vEBs")
	}
}
