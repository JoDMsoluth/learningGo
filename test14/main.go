package main

import "fmt"

// slice는 하나의 struct로 볼 수 있다.
// 구성요소는
// {
//	 pointer * (시작주소)
//	 len 	int
//	 cap 	int
// }
//

type Student struct {
	name  string
	age   int
	grage int
}

func (t *Student) SetName(newName string) {
	t.name = newName
}

func main() {
	arr1 := []int{1, 2, 3}

	// 배열은 주소를 복사한다.
	arr2 := arr1
	arr2[1] = 0
	fmt.Println(arr1, arr2)

	// 구조체는 값을 복사한다.
	a := Student{"aaa", 20, 10}
	b := &a
	b.SetName("bbb")
	fmt.Println(a, b)
	fmt.Println()

	// Slice는 reference type이다.
	arr3 := arr1[:3]
	var arr4 []int
	arr4 = arr3
	arr4[1] = 0
	arr3[2] = 999
	fmt.Printf("%p %p %p %p \n", arr1, arr2, arr3, arr4)
	fmt.Println(arr1, arr2, arr3, arr4)
}

// Instance란, OOP에서 주체를 의미한다.
// 주체는 속성과 기능이 존재한다.
// Instance는 struct 같은 개체를 말한다. -> 생명을 가진 개체로 생각하겠다.
// Instance는 생명주기를 갖는다.
// a.SetName("bbb")에서 a는 주체, SetName은 기능, "bbb"는 목적이다.
