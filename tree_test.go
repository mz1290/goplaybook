package main

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	root := createTree([]int{1, 2, 3, 4, 5, 6, 7})
	preorder(root)
	fmt.Println()
	inorder(root)
	fmt.Println()
	postorder(root)
	fmt.Println()
}

func TestIterativeTraversal(t *testing.T) {
	root := createTree([]int{1, 2, 3, 4, 5, 6, 7})
	iterativePreorder(root)
	fmt.Println()
	iterativeInorder(root)
	fmt.Println()
}

func TestLevelOrderTraversal(t *testing.T) {
	root := createTree([]int{1, 2, 3, 4, 5, 6, 7})
	levelorder(root)
	fmt.Println()
}

func TestNodeCountTreeHeight(t *testing.T) {
	root := createTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(countNodes(root))
	fmt.Println(treeHeight(root))
}

func TestBST(t *testing.T) {
	root := insertBST(nil, 10)
	root = insertBST(root, 5)
	root = insertBST(root, 20)
	root = insertBST(root, 8)
	root = insertBST(root, 30)

	deleteBSTR(root, 20)

	inorder(root)
	fmt.Println()

	found := searchBST(root, 20)
	if found != nil {
		fmt.Printf("found: %d\n", found.Data)
	}
}

func TestBST2(t *testing.T) {
	root := insertBST(nil, 50)
	root = insertBST(root, 10)
	root = insertBST(root, 40)
	root = insertBST(root, 20)
	root = insertBST(root, 30)

	deleteBSTR(root, 50)

	inorder(root)
	fmt.Println()

	found := searchBST(root, 20)
	if found != nil {
		fmt.Printf("found: %d\n", found.Data)
	}
}
