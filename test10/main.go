package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 숫자 야구

/*
	컴퓨터 : 무작위 숫자 3개
	-> 사용자 입력
	-> 비교
	-> 결과도출
*/

type Result struct {
	strikes int
	balls   int
}

func main() {
	// 항상 변하는 값을 만들기 위해 seed를 넣어준다.
	rand.Seed(time.Now().UnixNano())
	// 무작위 숫자 3개를 만든다.
	numbers := MakeNumbers()

	cnt := 0
	for {
		// 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3S 인가?
		if IsGameEnd(result) {
			// 게임 끝
			break
		}
		cnt++
	}

	// 게임 끝, 몇 번만에 맞췄는지 출력
	fmt.Printf("%d번 만에 맞췄습니다.\n", cnt)
}

func MakeNumbers() [3]int {
	// 0~9까지 무작이 숫자 3개를 만든다.
	var rst = [3]int{}

	// 겹치지 않는 3개의 숫자를 뽑는다.
	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}
	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	// 키보드로부터 0, 9 사이의 겹치지 않는 숫자를 입력하여 반환한다.
	var rst [3]int
	for {
		fmt.Println("겹치지 않는 0~9 숫자를 3개 입력해주세요")
		var no int
		// scanf할 때 무조건 '\n'을 포함시켜야 한다. 버퍼에 맞추기 위해
		_, err := fmt.Scanf("%d\n", &no)
		// nil 아무것도 아닌경우 (즉, 에러가 없다.)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}
		if no < 100 || no >= 1000 {
			fmt.Println("3개의 숫자를 입력해야 합니다.")
			continue
		}
		success := true
		idx := 0
		for no > 0 {
			// 입력값이 123 이면 [1, 2, 3]
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}
		if !success {
			continue
		}
		rst[0], rst[2] = rst[2], rst[0]
		break
	}
	return rst

}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	// 두개의 숫자 3개를 비교해서 결과를 반환한다.
	strikes := 0
	balls := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
			}
		}
	}
	return Result{strikes, balls}
}

func PrintResult(result Result) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)
}

func IsGameEnd(result Result) bool {
	// 결과를 입력받아 게임이 끝났는지 반환
	return result.strikes == 3
}
