package dataStruct

import "fmt"

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	// 현재 값 보다 작은 경우 왼쪽에 넣는다.
	if n.Val > v {
		if n.Left == nil {
			n.Left = &BinaryTreeNode{Val: v}
			return n.Left
		} else {
			return n.Left.AddNode(v)
		}
	} else {
		// 현재 값 보다 큰 경우 오른쪽에 넣는다.
		if n.Right == nil {
			n.Right = &BinaryTreeNode{Val: v}
			return n.Right
		} else {
			return n.Right.AddNode(v)
		}
	}
}

type DepthNode struct {
	depth int
	node  *BinaryTreeNode
}

func (t *BinaryTree) Print() {
	q := []DepthNode{}
	q = append(q, DepthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(q) > 0 {
		var first DepthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}
		fmt.Print(first.node.Val, " ")
		if first.node.Left != nil {
			q = append(q, DepthNode{depth: currentDepth + 1, node: first.node.Left})
		}
		if first.node.Right != nil {
			q = append(q, DepthNode{depth: currentDepth + 1, node: first.node.Right})
		}
	}
}

func (t *BinaryTree) Search(v int) (bool, int) {
	return t.Root.Search(v, 1)
}

func (n *BinaryTreeNode) Search(v int, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	} else if n.Val > v {
		if n.Left != nil {
			return n.Left.Search(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.Right != nil {
			return n.Right.Search(v, cnt+1)
		}
		return false, cnt
	}
}
