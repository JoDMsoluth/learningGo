package test18

import (
	"fmt"

	"github.com/JoDMsoluth/learningGo/dataStruct"
)
// DFS 만드는 법
// 1. 재귀호출
// 2. 스택

func Test18() {
	tree := dataStruct.Tree()
	val := 1
	tree.AddNode(1)
	val++

	// 루트 노드에 1, 2, 3 자식 노드추가
	for i:=-; i<3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	// 각 노드의 자식을 2개씩 추가
	for i:=0; i<len(tree.Root.Childs); i++ {
		for j:=0; j<2; i++ {
			tree.Root.Chilids[i].AddNode(val)
			val++
		}
	}

	// 1. 재귀호출을 이용한 DFS 순회
	tree.DFS1()
	fmt.Println()
	// 2. 스택을 이용한 DFS 순회
	// 2-1 루트를 스택에 넣는다.
	// 루트를 스택에서 출력
	// 2-2 루트의 자식들을 스택에 넣는다.
	// 루트의 자식들을 스택에서 출력
	// 2-3 루트의 첫번째 자식의 자식들을 스택에 넣는다. .....
	tree.DFS2()

	tree.BFS1()

}
