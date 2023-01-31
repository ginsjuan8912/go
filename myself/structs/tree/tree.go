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
	Root        *BinaryNode
	restriction int64
}

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

func (t *BinaryTree) Initialize(data int64) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		err := t.Root.Insert(data)
		switch err.(type) {
		case EmptyInsertionMessage:
			fmt.Println(err, "Cannot insert an empty number")
		}
	}

	return t
}

func (n *BinaryNode) Insert(data int64) error {

	if n == nil {
		return EmptyInsertionMessage{}
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			err := n.left.Insert(data)
			if err != nil {
				return err
			}
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			err := n.right.Insert(data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func PrintTree(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	PrintTree(w, node.left, ns+2, 'L')
	PrintTree(w, node.right, ns+2, 'R')
}
