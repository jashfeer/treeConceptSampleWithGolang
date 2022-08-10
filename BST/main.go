package main

import "fmt"

type Node struct {
	key   int
	Left  *Node
	Right *Node
}
func (n *Node) Insert(k int) {
	if n.key < k {
		if n.Right == nil {
			n.Right = &Node{key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.key > k {
		if n.Left == nil {
			n.Left = &Node{key: k}
		} else {
			n.Left.Insert(k)
		}

	}
}
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.key < k {
		return n.Right.Search(k)
	} else if n.key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{key: 100}

	fmt.Println(tree)
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(76))
	fmt.Println(tree.Search(8))
}
