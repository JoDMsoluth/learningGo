package test21

import (
	"fmt"

	"github.com/JoDMsoluth/learningGo/dataStruct"
)

// Map 과 Hash의 관계
// Hash 특정
// 1. 출력값의 범위가 있다. (ex 0 ~ 99)
// 2. 같은 입력이면 같은 출력
// 3. 다른 입력이면 다른 출력
// 4. one-way function 이다

// Rolling Hash
// "abcde" ---> Hash ---> "H0 ~ H4"
// S0 ~ Sn
// A = 256 : ASCII CODE
// B 소수 중에서 아무거나
// Hi = (H(i-1) * A + Si) mod B

// Hash 입력값이 크더라도 범위내에 값이 나온다. (손실압축)
// 충돌이 일어날 수 있다.
// 성능이 Hash func에 따라 다르다. 보통 O(1)
// hash가 순서에 상관없이 저장되어 있어 정렬된 형태로 값을 뽑아낼 수 없다.
// -> 보완 ) Sorted Map

func Test21() {
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("tbcde = ", dataStruct.Hash("tbcde"))
	fmt.Println("abcdf = ", dataStruct.Hash("abcdf"))
	fmt.Println("sdadsadfw = ", dataStruct.Hash("sdadsadfw"))

	m := dataStruct.CreateMap()
	m.Add("AAA", "01077777777")
	m.Add("BBB", "01077534237")
	m.Add("CCC", "01023252167")
	m.Add("DDD", "01016734587")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("DDD = ", m.Get("DDD"))
	fmt.Println("EEE = ", m.Get("EEE"))

	// Go에서 제공하는 Map
	// 선언
	var m1 map[string]string
	// 반드시 초기화 해서 사용
	// 초기화
	m1 = make(map[string]string)

	m1["bcd"] = "ccc"
	m1["abc"] = "bbb"
	fmt.Println("m1['abc'] = ", m1["abc"])

	m2 := make(map[int]int)
	m2[51] = 23
	fmt.Println("m2[51] = ", m2[51])
	fmt.Println("m2[53] = ", m2[53])

	m3 := make(map[int]bool)
	m3[4] = true
	fmt.Println(m3[6], m3[4])

	// ok를 통해 키값이 존재하는지 않하는지 확인
	v, ok := m2[51]
	v1, ok1 := m2[53]
	fmt.Println(v, ok, v1, ok1)

	// 삭제
	delete(m2, 51)
	v2, ok2 := m2[51]
	fmt.Println(v2, ok2)

	// 순회
	m2[61] = 23
	m2[31] = 153
	m2[41] = 235
	for key, value := range m2 {
		fmt.Println("key : ", key, "value : ", value)
	}

}
