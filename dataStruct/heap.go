package dataStruct

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)
	// 맨 끝 노드 인덱싱
	idx := len(h.list) - 1
	for idx >= 0 {
		// 부모 인덱스
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		// 맨 끝과 맨 끝의 부모 스와핑
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}

}

func (h *Heap) Pop() int {
	// 빈경우
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	// 맨 뒤에 거 제고하고 리스트를 뺀후
	h.list = h.list[:len(h.list)-1]

	// 맨 뒤에 있는걸 맨 위로 올린다.
	h.list[0] = last
	idx := 0

	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1
		// 비교할 자식 노드가 없는 경우
		if leftIdx >= len(h.list) {
			break
		}

		// 왼쪽 자식이 현재 값보다 크다
		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			// 오른쪽 자식이 현재 값보다 크다
			if h.list[rightIdx] > h.list[idx] {
				// 왼쪽이랑 비교해서 오른쪽이랑 왼쪽중 큰 쪽이랑 부모랑 바꾼다.
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}
		// 바꿀 자식노드가 없다(둘다 큰 조건 만족하지 않는다)
		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}
