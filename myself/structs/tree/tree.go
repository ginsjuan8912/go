package tree

import (
	"fmt"
	"io"
)

type RestrictionMessage struct{}

type EmptyInsertionMessage struct{}

func (m RestrictionMessage) Error() string {
	return "The number cannot be inserted it fails the restriction"
}

func (m EmptyInsertionMessage) Error() string {
	return "Cannot insert empty node"
}

type BinaryTree struct {
	root        *BinaryNode
	restriction int64
}

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		err := t.root.insert(data)
		switch err.(type) {
		case EmptyInsertionMessage:
			fmt.Println(err, "Cannot insert an empty number")
		}
	}

	return t
}

func (n *BinaryNode) insert(data int64) error {

	if n == nil {
		return EmptyInsertionMessage{}
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			err := n.left.insert(data)
			if err != nil {
				return err
			}
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			err := n.right.insert(data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func printTree(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	printTree(w, node.left, ns+2, 'L')
	printTree(w, node.right, ns+2, 'R')
}
