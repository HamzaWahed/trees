package trees

import (
	"fmt"
	"testing"
)

func initializeSmallCompleteTree() (*SplayTree, int32) {
	tree := NewSplayTree(3)
	var lastElement int32 = 0
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(lastElement)
	return tree, lastElement
}

func initializeMediumTree() *SplayTree {
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
	tree, _ := initializeSmallCompleteTree()
	tree.Search(1)
	if tree.Root.Data != 1 {
		t.Errorf("Tree root is %d, expected 1", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_ElementNotInTree(t *testing.T) {
	tree, _ := initializeSmallCompleteTree()
	tree.Search(-1)
	if tree.Root.Data != 0 {
		t.Errorf("Tree root is %d, expected 0", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag(t *testing.T) {
	tree, _ := initializeSmallCompleteTree()
	tree.Search(5)
	if tree.Root.Data != 5 {
		t.Errorf("Tree root is %d, expected 5", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag_ElementNotInTree(t *testing.T) {
	tree, _ := initializeSmallCompleteTree()
	tree.Search(10)
	if tree.Root.Data != 6 {
		t.Errorf("Tree root is %d, expected 6", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Zig(t *testing.T) {
	tree, _ := initializeSmallCompleteTree()
	tree.Search(6)
	if tree.Root.Data != 6 {
		t.Errorf("Tree root is %d, expected 6", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Zag(t *testing.T) {
	tree, _ := initializeSmallCompleteTree()
	tree.Search(2)
	tree.PrintTree()
	if tree.Root.Data != 2 {
		t.Errorf("Tree root is %d, expected 2", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag_Zig(t *testing.T) {
	tree, _ := initializeSmallCompleteTree()
	tree.Search(4)
	tree.PrintTree()
	if tree.Root.Data != 4 {
		t.Errorf("Tree root is %d, expected 4", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Medium_Tree(t *testing.T) {
	tree := initializeMediumTree()
	tree.Search(5)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 5 {
		t.Errorf("Tree root is %d, expected 5", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zag_Medium_Tree(t *testing.T) {
	tree := initializeMediumTree()
	tree.Search(15)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 15 {
		t.Errorf("Tree root is %d, expected 15", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Zig_Medium_Tree(t *testing.T) {
	tree := initializeMediumTree()
	tree.Search(0)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 0 {
		t.Errorf("Tree root is %d, expected 0", tree.Root.Data)
	}
}

func TestSplayTree_Search_Zig_Zag_Medium_Tree(t *testing.T) {
	tree := initializeMediumTree()
	tree.Search(2)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 2 {
		t.Errorf("Tree root is %d, expected 2", tree.Root.Data)
	}
}

func TestSplayTree_Insert_Small_Tree1(t *testing.T) {
	tree, root := initializeSmallCompleteTree()
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != root {
		t.Errorf("Tree root is %d, expected %d", tree.Root.Data, root)
	}
}

func TestSplayTree_Insert_Small_Tree2(t *testing.T) {
	tree, _ := initializeSmallCompleteTree()
	tree.Insert(20)
	tree.PrintTree()
	fmt.Println()
	if tree.Root.Data != 20 {
		t.Errorf("Tree root is %d, expected 20", tree.Root.Data)
	}
}
