package main

import "github.com/joDmsoluth/learningGo/dataStruct"

func main() {
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
