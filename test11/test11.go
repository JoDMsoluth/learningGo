package test11

import "fmt"

// 메모리를 할당하는 구조 -> 8byte가 필요하면 컴퓨터가 8byte를 준다.
// 그외에는 다른 정보가 들어있기 때문에 마음대로 공간을 늘릴 수 없다.
// 동적배열은 이미 할당한 배열 말고 다른 공간에 또 다른 배열을 만들어 할당하고 기존 배열을 복사해서 사용한다.
// 그 후 새로만든 배열을 포인팅 한다.

// slice : 동적배열 - 실제 배열을 포인트하고 있다.
// 용어1 : capacity : 컴퓨터가 할당한 공간
// 용어2 : length : 사용하고 있는 공간

func Test11() {
	// 동적배열 사용
	a := []int{1, 2, 3}
	fmt.Printf("len(a) = %d\n", len(a)) // 3
	fmt.Printf("cap(a) = %d\n", cap(a)) // 3

	b := make([]int, 0, 8)
	fmt.Printf("len(a) = %d\n", len(b)) // 0
	fmt.Printf("cap(a) = %d\n", cap(b)) // 8

	// Element 추가 방법
	c := append(a, 4)
	fmt.Printf("len(a) = %d\n", len(c)) // 4
	fmt.Printf("cap(a) = %d\n", cap(c)) // 6 : 더 추가할 것을 생각하여 2배로 늘려서 만든다.

	// 여유공간(cap)이 없을 때 주소가 바뀌었는가?
	// 결론 : 바뀜
	fmt.Printf("이전 주소 : %p, 이후 주소 : %p\n", a, c)

	// 여유공간(cap)이 없을 때 주소가 바뀌었는가?
	// 결론 : 안바뀜
	d := append(c, 4)
	fmt.Printf("이전 주소 : %p, 이후 주소 %p\n", c, d)

	// 어쩔 때는 바뀌고 안바뀌고 어려우니 다음과 같은 패턴을 사용한다.
	// 즉, 복사해서 사용하자 -> 언제든 다른 주소로 나온다
	e := make([]int, 2, 4)
	e[0] = 1
	e[1] = 2

	// e크기만큼 만든다
	f := make([]int, len(e))

	// e를 복사한다.
	for i := 0; i < len(e); i++ {
		f[i] = e[i]
	}

	// 무조건 다른 메모리를 반환한다.
	f = append(f, 3)

	fmt.Printf("%p %p\n", e, f)
	f[0] = 4
	f[1] = 5

	fmt.Println(e)
	fmt.Println(f)
}
