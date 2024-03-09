package trees

import (
	"fmt"
	"testing"
)

func intializeSmallCompleteTree() *SplayTree {
	tree := NewSplayTree(3)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(0)
	return tree
}

func intializeMediumTree() *SplayTree {
	tree := NewSplayTree(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(12)
	tree.Insert(20)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(9)
	tree.Insert(0)
	tree.Insert(2)
	tree.Insert(6)
	fmt.Println("Medium Tree initialized:")
	tree.PrintTree()
	fmt.Println()
	return tree
}

func TestSplayTree_Search_Zig(t *testing.T) {
	tree := intializeSmallCompleteTree()
	tree.Search(1)
	if tree.Root.Data != 1 {
		t.Errorf("Tree root is %d, expected 1", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_ElementNotInTree(t *testing.T) {
	tree := intializeSmallCompleteTree()
	tree.Search(-1)
	if tree.Root.Data != 0 {
		t.Errorf("Tree root is %d, expected 0", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag(t *testing.T) {
	tree := intializeSmallCompleteTree()
	tree.Search(5)
	if tree.Root.Data != 5 {
		t.Errorf("Tree root is %d, expected 5", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag_ElementNotInTree(t *testing.T) {
	tree := intializeSmallCompleteTree()
	tree.Search(10)
	if tree.Root.Data != 6 {
		t.Errorf("Tree root is %d, expected 6", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Zig(t *testing.T) {
	tree := intializeSmallCompleteTree()
	tree.Search(6)
	if tree.Root.Data != 6 {
		t.Errorf("Tree root is %d, expected 6", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Zag(t *testing.T) {
	tree := intializeSmallCompleteTree()
	tree.Search(2)
	tree.PrintTree()
	if tree.Root.Data != 2 {
		t.Errorf("Tree root is %d, expected 2", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag_Zig(t *testing.T) {
	tree := intializeSmallCompleteTree()
	tree.Search(4)
	tree.PrintTree()
	if tree.Root.Data != 4 {
		t.Errorf("Tree root is %d, expected 4", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Medium_Tree(t *testing.T) {
	tree := intializeMediumTree()
	tree.Search(5)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 5 {
		t.Errorf("Tree root is %d, expected 5", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag_Medium_Tree(t *testing.T) {
	tree := intializeMediumTree()
	tree.Search(15)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 15 {
		t.Errorf("Tree root is %d, expected 15", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Zig_Medium_tree(t *testing.T) {
	tree := intializeMediumTree()
	tree.Search(0)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 0 {
		t.Errorf("Tree root is %d, expected 0", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Zag_Medium_tree(t *testing.T) {
	tree := intializeMediumTree()
	tree.Search(2)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 2 {
		t.Errorf("Tree root is %d, expected 2", tree.Root.Data)
	}
}

func TestSplayTree_Insert_Zig(t *testing.T) {

}

func TestSplayTree_PrintTree(t *testing.T) {

}
