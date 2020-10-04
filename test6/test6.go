package test6

import "fmt"

func add(x int, y int) int {
	return x + y
}

func Test6() {

	// 반복문
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	// 재귀
	s := sum(10, 0)
	fmt.Println(s)

	// 피보나치 수열
	f := f(3)
	fmt.Println(f)
}

// 모든 재귀호출은 반복문으로 바꿀 수 있다.
// 재귀호출이 쉬운 경우 ex ) 피포나치 수열
// 수학의 정의를 프로그램으로 만들 때 재귀호출이 쉬운 경우가 많다.
// 그 외에는 반복문이 대채로 쉽다.
// 반면 재귀호출은 함수 호출과정이 어렵고 단계가 많다.(느리다)
func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}

func f(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return f(x-1) + f(x-2)
}
