package test8

import "fmt"

// First Class : go 언어는 속성뿐 아니라 기능도 가지고 있따.
// 반면 C언어는 속성만 가지고 있다.
type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

// 기능
func (s Student) ViewSungJuk() {
	fmt.Println(s.sungjuk)
}
func (s *Student) InputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

// 비교 ) 그냥 일반함수
// func ViewSungJuk() (s Student) {
// 	fmt.Println(s.sungjuk)
// }

func Test8() {
	var s Student
	s.name = "철수"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"

	fmt.Println(s)

	s.InputSungjuk("과학", "B")
	// "수학" ,"C"로 출력되는 이유
	// Go에서 함수의 인자는 무조건 복사만 이루어진다.
	// 떄문에 받을 떄 주소로 받아야된다.
	fmt.Println(s)
}
