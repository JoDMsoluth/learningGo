package test16

import (
	"github.com/JoDMsoluth/learningGo/dataStruct"
)

// double linkedList.go
func Test16() {
	list := &dataStruct.DoubleLinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNode()
	list.RemoveNode(list.Root)
	list.PrintNode()
	list.RemoveNode(list.Root.Next)
	list.PrintNode()
	list.RemoveNode(list.Tail)
	list.PrintNode()

	list.PrintReverse()
}
