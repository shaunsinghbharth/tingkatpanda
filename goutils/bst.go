package goutils

import (
	"fmt"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type BSTNode struct {
	Value string
	Data  string
	Left  *BSTNode
	Right *BSTNode
	bal   int
}

func (n *BSTNode) Insert(value, data string) bool {

	switch {
	case value == n.Value:
		n.Data = data
		return false
	case value < n.Value:

		if n.Left == nil {

			n.Left = &BSTNode{Value: value, Data: data}

			if n.Right == nil {

				n.bal = -1
			} else {

				n.bal = 0
			}
		} else {

			if n.Left.Insert(value, data) {

				if n.Left.bal < -1 || n.Left.bal > 1 {
					n.rebalance(n.Left)
				} else {

					n.bal--
				}
			}
		}

	case value > n.Value:
		if n.Right == nil {
			n.Right = &BSTNode{Value: value, Data: data}
			if n.Left == nil {
				n.bal = 1
			} else {
				n.bal = 0
			}
		} else {
			if n.Right.Insert(value, data) {
				if n.Right.bal < -1 || n.Right.bal > 1 {
					n.rebalance(n.Right)
				} else {
					n.bal++
				}
			}
		}
	}
	if n.bal != 0 {
		return true
	}

	return false
}

func (n *BSTNode) rotateLeft(c *BSTNode) {

	r := c.Right

	c.Right = r.Left

	r.Left = c

	if c == n.Left {
		n.Left = r
	} else {
		n.Right = r
	}

	c.bal = 0
	r.bal = 0
}

func (n *BSTNode) rotateRight(c *BSTNode) {

	l := c.Left
	c.Left = l.Right
	l.Right = c
	if c == n.Left {
		n.Left = l
	} else {
		n.Right = l
	}
	c.bal = 0
	l.bal = 0
}

func (n *BSTNode) rotateRightLeft(c *BSTNode) {

	c.Right.Left.bal = 1
	c.rotateRight(c.Right)
	c.Right.bal = 1
	n.rotateLeft(c)
}

func (n *BSTNode) rotateLeftRight(c *BSTNode) {
	c.Left.Right.bal = -1
	c.rotateLeft(c.Left)
	c.Left.bal = -1
	n.rotateRight(c)
}

func (n *BSTNode) rebalance(c *BSTNode) {

	switch {

	case c.bal == -2 && c.Left.bal == -1:
		n.rotateRight(c)

	case c.bal == 2 && c.Right.bal == 1:
		n.rotateLeft(c)

	case c.bal == -2 && c.Left.bal == 1:
		n.rotateLeftRight(c)

	case c.bal == 2 && c.Right.bal == -1:
		n.rotateRightLeft(c)
	}
}

func (n *BSTNode) Find(s string) (string, bool) {

	if n == nil {
		return "", false
	}

	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

func (n *BSTNode) Dump(i int, lr string) {
	if n == nil {
		return
	}
	indent := ""
	if i > 0 {

		indent = strings.Repeat(" ", (i-1)*4) + "+" + lr + "--"
	}
	fmt.Printf("%s%s[%d]\n", indent, n.Value, n.bal)
	n.Left.Dump(i+1, "L")
	n.Right.Dump(i+1, "R")
}

type BinarySearchTree struct {
	Root *BSTNode
}

func (t *BinarySearchTree) Insert(value, data string) {
	if t.Root == nil {
		t.Root = &BSTNode{Value: value, Data: data}
		return
	}
	t.Root.Insert(value, data)

	if t.Root.bal < -1 || t.Root.bal > 1 {
		t.rebalance()
	}
}

func (t *BinarySearchTree) rebalance() {
	fakeParent := &BSTNode{Left: t.Root, Value: "fakeParent"}
	fakeParent.rebalance(t.Root)

	t.Root = fakeParent.Left
}

func (t *BinarySearchTree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

func (t *BinarySearchTree) Traverse(n *BSTNode, f func(*BSTNode)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func (t *BinarySearchTree) Dump() {
	t.Root.Dump(0, "")
}
