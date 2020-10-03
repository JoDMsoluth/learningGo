package main

import (
	// 버퍼에 관한 거
	"fmt"
	// 문자를 숫자로 바꿔주는거
	// 문자열에 대한 거
)

func main() {
	var i int
	for i < 10 {
		fmt.Println(i)
		i += 1
	}

	// scope에 의해 새로 생성된 i
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for true {
		break
	}

	// 구구단
	for dan := 1; dan <= 9; dan++ {
		fmt.Printf("%d단\n", dan)
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
		}
	}

	// 구구단
	for i := 1; i <= 9; i++ {
		fmt.Printf("\n")
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
	}
}
