package test15

import "fmt"

// 링크드 리스트

type Node struct {
	next *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) PrintNode() {
	node := l.root
	for node != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Println()
}

func (l *LinkedList) AddNode(val int) {
	// 처음 아무것도 없을 떄, root, tail은 nil이다.
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(node *Node) {
	// 루트를 지울 때
	if node == l.root {
		l.root = l.root.next
		node.next = nil
		return
	}

	// 지우고자 하는 노드가 나올 때까지 전진한다.
	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	// tail을 지울 때
	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		// 중간 노드 지울 때
		prev.next = prev.next.next
	}
	node.next = nil
}

func Test15() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNode()
	list.RemoveNode(list.root)
	list.PrintNode()
	list.RemoveNode(list.root.next)
	list.PrintNode()
	list.RemoveNode(list.tail)
	list.PrintNode()

	fmt.Println("/////////////////////////////////////////////////////////////////")
	var root *Node
	var tail *Node

	// 처음엔 tail과 root가 같다. 만약 추가되면 거기가 tail이다.
	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	node := root

	PrintNode(node)

	node = root
	// root 다음 노드를 뺀다.
	root, tail = RemoveNode(root.next, root, tail)
	fmt.Printf("root : %d, tail : %d\n", root.val, tail.val)
	PrintNode(node)
	root, tail = RemoveNode(tail, root, tail)
	fmt.Printf("root : %d, tail : %d\n", root.val, tail.val)
	PrintNode(node)
}

func PrintNode(node *Node) {
	// 끝까지 갈때까지 반복문으로 돌림
	for node != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Println()
}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}

	// 마지막에 노드를 추가
	tail.next = node
	return node
}

// 지우고자 하는 노드의 이전 노드가 필요하다
// 그럼 맨 앞에 노드는 어떻게? -> root를 그냥 다음 노드로 바꿔준다.
// 그럼 맨 끝의 노드는 어떻게? -> tail을 그냥 그전 노드로 바꿔준다.
// 그럼 노드가 하나 밖에 없을 떄는 어떻게? -> root와 tail이 nil을 가리킨다.
func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	// 지우고자 하는 게 root다.
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}
	// 지우고자 하는 게 root가 아니다.
	prev := root
	// node 이전 노드까지 위치시킨다.
	for prev.next != node {
		prev = prev.next
	}

	// 지우고자 하는게 tail이다.
	// prev의 next를 지우고
	// prev가 tail이 된다.
	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		// 중간이면 다다음 노드와 연결시킨다.
		prev.next = prev.next.next
	}
	return root, tail
}
