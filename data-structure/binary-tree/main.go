package main

import "fmt"

// Node sdfsdfsdf
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) insert(k int) {
	if n.Key < k {
		//right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.insert(k)
		}
	} else if n.Key > k {
		//left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.insert(k)
		}
	}

}

func (n *Node) print() {
	fmt.Println(n.Key)

	if n.Left != nil {
		n.Left.print()
	}

	if n.Right != nil {
		n.Right.print()
	}
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	//tree.Insert(200)
	//tree.Insert(300)
	//fmt.Println(tree)

	tree.insert(52)
	tree.insert(203)
	tree.insert(19)
	tree.insert(76)
	tree.insert(310)
	tree.insert(7)

	tree.print()

}
