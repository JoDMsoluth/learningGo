package test25

import "fmt"

// 절차지향 프로그래밍
// 장점 : 메인 함수 부분이 간단하고 보기 쉽다.
// 단점 : 재사용이 힘들고 함수 부분이 지저분하다
// 치명적 단점 : 딸기잼이 오랜지 잼으로 바뀌면 기존 func을 그대로 쓸수 없고 다시 만들어야 한다.

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func GetOneSpoon(jam *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}

func Test25() {
	// 1. 빵 두개를 꺼낸다.
	breads := GetBreads(2)

	jam := &StrawberryJam{}

	// 2. 딸기잼 뚜껑을 연다.
	OpenStrawberryJam(jam)

	// 3. 딸기잼 한 스푼 뜬다.
	spoon := GetOneSpoon(jam)

	// 4. 딸기잼을 빵에 바른다.
	PutJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)

}
