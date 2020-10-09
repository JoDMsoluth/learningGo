package test26

import "fmt"

// Object는 상태 + 기능
// interface는 Object에서 기능만을 모아 정의한 것이다.

// interface는 관계만 선언, 공통된 기능들을 묶어 준다.
// 관계를 분리해놓음으로서 object의 종속성을 풀고 확장성을 올려준다.

type SpoonOfJam interface {
	String() string
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

type StrawberryJam struct{}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

type OrangeJam struct{}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type AppleJam struct{}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type SpoonOfStrawberryJam struct{}

type SpoonOfOrangeJam struct{}
type SpoonOfAppleJam struct{}

func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberry"
}
func (s *SpoonOfOrangeJam) String() string {
	return " + orange"
}
func (s *SpoonOfAppleJam) String() string {
	return " + apple"
}

// Bread입장에서는 StrawberryJam이든 OrangeJam이든 Jam에 대한 관계만 가지고 있으면 된다.
func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread + " + b.val
}

func Test26() {
	bread := &Bread{val: "cheese"}
	// jam := &StrawberryJam{}
	// jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
