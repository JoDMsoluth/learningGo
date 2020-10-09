package test19

import (
	"fmt"

	"github.com/JoDMsoluth/learningGo/dataStruct"
)

// 이진 탐색 트리 BST
// 왼쪽에는 부모보다 작은 숫자 오른쪽은 부모보다 큰 숫자가 오도록 한다.
//    4
//  2	5
// 1 3	 6
// 이 트리에서 3을 찾을 때 4(루트)의 왼쪽 2의 오른쪽을 찾아면 된다.
// 즉, 어떤 값을 찾을 때, 전부 순회할 필요 없이, logN 의 효율이 나온다.
// 최소신장트리를 만들어야 탐색 효율이 나온다.
// 최소신장트리 만드는 법 (AVL 트리, Black Red 트리)

func Test19() {
	tree := dataStruct.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()
	fmt.Println()

	if found, cnt := tree.Search(6); found {
		fmt.Println("count : ", cnt, " found : 11")
	} else {
		fmt.Println("count : ", cnt, " not found : 11")
	}

}
