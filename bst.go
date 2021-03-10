package goutils

import "fmt"

type BinaryNode struct {
	item  string
	left  *BinaryNode
	right *BinaryNode
}

type BST struct {
	root *BinaryNode
}

func (bst *BST) insertNode(t **BinaryNode, item string) error {
	if *t == nil {
		newNode := &BinaryNode{
			item:  item,
			left:  nil,
			right: nil,
		}
		*t = newNode

		return nil
	}

	if item < (*t).item {
		bst.insertNode(&((*t).left), item)
	} else {
		bst.insertNode(&((*t).right), item)
	}
	return nil
}

func (bst *BST) Insert(item string) {
	bst.insertNode(&bst.root, item)
}

func (bst *BST) inOrderTraverse(t *BinaryNode) {
	if t != nil {
		bst.inOrderTraverse(t.left)
		fmt.Println(t.item)
		bst.inOrderTraverse(t.right)
	}
}

func (bst *BST) InOrder() {
	bst.inOrderTraverse(bst.root)
}

func (bst *BST) Search(match string) {
	bst.searchInOrderTraverse(bst.root, match)
}

func (bst *BST) searchInOrderTraverse(t *BinaryNode, match string) {
	if t != nil {
		bst.inOrderTraverse(t.left)
		if match == t.item {
			fmt.Println(t.item)
			return
		}
		bst.inOrderTraverse(t.right)
	}
}

func (bst *BST) Size() int {
	return bst.Count(bst.root)
}

func (bst *BST) Count(t *BinaryNode) int {
	if t == nil {
		return 0
	} else {
		return bst.Count(t.left) + 1 + bst.Count(t.right)
	}
}
