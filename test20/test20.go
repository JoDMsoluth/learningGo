package test20

import (
	"fmt"

	"github.com/JoDMsoluth/learningGo/dataStruct"
)

// 힙
// 최대힙 : 루트가 가장 큰 수가 와야 됨, 부모 노드보다 작은 수가 와야 됨
// 최소힙 : 루트가 가장 작은 수가 와야 됨. 부모 노드보다 큰 수가 와야 됨
// 만드는 법

// 1. push : 새로 붙인 애는 맨 마지막에 놓는다.
// -> 그후 부모랑 비교 ->
// 부모보다 크면 서로 바꾼다. ->
// 다시 개를 기준으로 또 부모랑 비교해서 크면 서로 바꾼다.
// -> 계속 반복

// 2. pop : 루트를 뺀다
// -> 루트의 빈자리를 맨 끝에 있는 애로 채운다.
// -> 자식 노드 중에 루트(부모)보다 큰 애가 있으면 그 차이가 가장 큰 애와 서로 바꾼다.
// -> 계속 반복

// 정렬로 활용 가능 -> pop하면 큰순 or 작은순으로 나온다.
// 속도 log2(N)

// 구현방법
// 1. 배열
// N번째 노드의 왼쪽자식 : 2N+1, 오른쪽 자식 : 2ㅜ+2
// N번째 노드의 부모 : (N-1)/2
// 2. List

func Test20() {
	h := dataStruct.Heap{}
	h.Push(9)
	h.Push(7)
	h.Push(5)
	h.Push(8)
	h.Push(6)
	h.Push(4)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
