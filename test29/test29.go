package test29

import "fmt"

// TDD
// 기존 : 설계 - 구현 - test
// TDD : 테스트 코드(실패) - 일단 되도록 코딩 -> 성공 강화 (리팩토링) -> 테스트 코드
// 이게 왜 좋은가??????

// 장점
// 1. Test case가 많아진다.
// 2. 촘촘한 test로 유지보수가 쉬워진다. - TDD를 하는 이유
// 3. 눈앞에 보이는 문제를 먼저 해결하는 코딩 패러다임 (재밌음)

// 단점
// 1. 유저스토리(유저 시나리오) 기반의 테스트 힘듦 (접속 -> 스크롤 -> 좋아요)
// 이유 ) TDD는 Unit test이기 때문
// 2. 흐지부지 되기가 쉽다. 관리하는 사람이 있어야 하고 팀원 전부가 할 줄 알아야함
// 해결 ) 관리하는 사람이 유지해야한다. 코드리뷰해야한다.
// 3. 멀티쓰레딩 환경에서 테스트하기 힘들다.
// 이유 ) 외부 모듈을 거쳐야 하는 상황에서 테스트가 힘들다.
// -> 응답을 잘 처리하고 싶은지 테스트 하려면
// -> 인증서버를 목업하고 테스트 한다.
// -> 그런데 목업 만들기가 귀찮다.

var opMap map[string]func(int, int) int

func initOpMap() {
	opMap = make(map[string]func(int, int) int)
	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mul
	opMap["/"] = div
	opMap["**"] = pow
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func pow(a, b int) int {
	rst := 1
	for i := 0; i < b; i++ {
		rst *= a
	}
	return rst
}

func Calculate(op string, a, b int) int {
	// 개선 : if-else -> switch-case -> Map
	if v, ok := opMap[op]; ok {
		return v(a, b)
	}
	return 0
}

func Test() {
	if !testCalculate("Test1", "+", 3, 2, 5) {
		return
	}

	if !testCalculate("Test2", "-", 5, 3, 2) {
		return
	}

	if !testCalculate("Test3", "-", 3, 6, -3) {
		return
	}

	if !testCalculate("Test4", "*", 3, 6, 18) {
		return
	}

	if !testCalculate("Test5", "/", 6, 3, 2) {
		return
	}

	// 추가 기능
	if !testCalculate("Test6", "**", 2, 3, 8) {
		return
	}

	fmt.Println("Success!")
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected: %d output: %d\n", testcase, expected, o)
		return false
	}
	return true
}

func Test29() {
	initOpMap()
	Test()
}
