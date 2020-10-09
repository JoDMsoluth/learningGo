package dataStruct

import "fmt"

// 더블 링크드 리스트
// 중간 노드를 가기 위해서 root부터 시작해서 찾았었는데
// 이제 중간에서 이전 노드를 찾아 처리할 수 있게 하기 위해 더블링크드리스트 사용

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

type DoubleLinkedList struct {
	Root *Node
	Tail *Node
}

func (l *DoubleLinkedList) PrintNode() {
	node := l.Root
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func (l *DoubleLinkedList) PrintReverse() {
	node := l.Tail

	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Println()
}

func (l *DoubleLinkedList) AddNode(val int) {
	// 처음 아무것도 없을 떄, Root, tail은 nil이다.
	if l.Root == nil {
		l.Root = &Node{Val: val}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{Val: val}
	prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev
}

func (l *DoubleLinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}
func (l *DoubleLinkedList) Front() int {
	if l.Root != nil {
		return l.Root.Val
	}
	return 0
}
func (l *DoubleLinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}
func (l *DoubleLinkedList) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}
func (l *DoubleLinkedList) Empty() bool {
	return l.Root == nil
}

func (l *DoubleLinkedList) RemoveNode(node *Node) {
	// 루트를 지울 때
	if node == l.Root {
		l.Root = l.Root.Next
		if l.Root != nil {
			l.Root.Prev = nil
		}
		node.Next = nil
		return
	}

	// 지우고자 하는 노드의 prev로 간다.
	Prev := node.Prev

	// tail을 지울 때
	if node == l.Tail {
		Prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = Prev
	} else {
		// 중간 노드 지울 때
		node.Prev = nil
		Prev.Next = Prev.Next.Next
		Prev.Next.Prev = Prev
	}
	node.Next = nil
}
