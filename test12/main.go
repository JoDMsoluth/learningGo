package main

import "fmt"

// 배열을 slice 짜른다. ex) a[start : end]
// start index 에서 시작해서 end index까지인데 end는 포함하지 않는다.
// slice는 새로 배열을 새로 만드는 것이 아니고 array의 부분은 가리키는 것이다.
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[4:8] // 5, 6, 7, 8
	c := a[4:]  // 5 ~ 10
	d := a[:4]  // 1 ~ 4

	// slice는 짤라서 새로운 slice를 만든게 아니다.
	// a의 부분적으로 가리키는 것이다.
	// 따라서 b가 바뀌면 a도 바뀐다.
	b[0] = 0
	b[1] = 0
	fmt.Println(b, c, d)
	fmt.Println(a)

	for i := 0; i < 5; i++ {
		var back int
		a, back = RemoveBack(a)
		fmt.Println("삭제 후:", a, ", pop된 값:", back)
	}
	fmt.Println(b, c, d)
	fmt.Println(a)

	for i := 0; i < 5; i++ {
		var front int
		a, front = FrontBack(a)
		fmt.Println("삭제 후:", a, ", pop된 값:", front)
	}
	fmt.Println(b, c, d)
	fmt.Println(a)
}

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func FrontBack(a []int) ([]int, int) {
	return a[1:], a[0]
}
