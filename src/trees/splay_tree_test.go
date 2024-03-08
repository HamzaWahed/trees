package trees

import "testing"

func TestSplayTree_Search_Zig(t *testing.T) {
	tree := NewSplayTree(3)
	tree.Insert(5)
	tree.Insert(1)
	tree.Search(1)
	if tree.Root.Data != 1 {
		t.Errorf("Tree root is %d, expected 1", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_ElementNotInTree(t *testing.T) {
	tree := NewSplayTree(3)
	tree.Insert(5)
	tree.Insert(1)
	tree.Search(2)
	if tree.Root.Data != 1 {
		t.Errorf("Tree root is %d, expected 1", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag(t *testing.T) {
	tree := NewSplayTree(3)
	tree.Insert(5)
	tree.Insert(1)
	tree.Search(5)
	if tree.Root.Data != 5 {
		t.Errorf("Tree root is %d, expected 5", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag_ElementNotInTree(t *testing.T) {
	tree := NewSplayTree(3)
	tree.Insert(5)
	tree.Insert(1)
	tree.Search(10)
	if tree.Root.Data != 5 {
		t.Errorf("Tree root is %d, expected 5", tree.Root.Data)
	}
}

func TestSplayTree_Insert_Zig(t *testing.T) {

}

func TestSplayTree_PrintTree(t *testing.T) {

}
