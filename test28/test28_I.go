package test28

// 이렇게 모은거 보다
type Actor interface {
	Move()
	Attact()
	Talk()
}

// 아래처럼 분리한 것이 더 좋다.
type Movable interface {
	Move()
}
type Attactable interface {
	Attact()
}
type Talkable interface {
	Talk()
}

// 모아 놓은다면 단일 책임 원칙에 반한다.
func MoveTo(a Actor) {
	a.Move()
	a.Attact()
}

// 분리해놓으면 강제로 move 기능만을 담당하게 할 수 있다.
func MoveTo(a Movable) {
	a.Move()
}
