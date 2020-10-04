package test9

import "fmt"

// 모든언어가 묵시적이든 명시적이든 포인터를 사용
// 포인터가 있으면 메모리 접근으로 보안상 위험하니 따로 관리하는 법을 배워야 한다.
// golang은 명시적으로 포인터를 사용하나 문제가 되는 연산이나 캐스팅의 사용을 막아놨다.

// 구조체의 경우 포인터를 호출하면 메모리 주소만(메모리 적게 사용)
// 값형태로 받으면 값 전체가 복사(메모리 많이 사용)
type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintSunguk() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSunguk(class, grade string) {
	s.class, s.grade = class, grade
}

func Test9() {
	var a int = 1
	var p *int

	Increase(&a)

	p = &a
	println(a, p, *p)

	var s Student = Student{name: "Tucker", age: 12, class: "과학", grade: "C"}

	s.InputSunguk("수학", "B")
	s.PrintSunguk()
}

func Increase(x *int) {
	*x++
}
