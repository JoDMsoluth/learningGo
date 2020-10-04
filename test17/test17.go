package test17

import (
	"fmt"

	"github.com/JoDMsoluth/learningGo/dataStruct"
)

// 스택, 큐
func Test17() {
	stack2 := dataStruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	fmt.Println("NewStack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)
	}

	queue2 := dataStruct.NewQueue()
	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Println("%d ->", val)
	}
}
