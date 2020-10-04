package test4

import (
	"bufio" // 버퍼에 관한 거
	"fmt"
	"os"      // os에 관한 거
	"strconv" // 문자를 숫자로 바꿔주는거
	"strings"
	// 문자열에 대한 거
)

func Test4() {
	fmt.Println("숫자를 입력하세요.")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n') // 한줄 읽어온다.
	line = strings.TrimSpace(line)     // 빈공간 제거

	n1, _ := strconv.Atoi(line) //Atoi A:문자 -> i:숫자

	line, _ = reader.ReadString('\n') // 한줄 읽어온다.
	line = strings.TrimSpace(line)    // 빈공간 제거

	n2, _ := strconv.Atoi(line) //Atoi A:문자 -> i:숫자

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.\n", n1, n2)

	fmt.Printf("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else if line == "*" {
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	} else {
		fmt.Printf("잘못 입력하셨습니다.")
	}

	// 만약 두가지가 해당된다면 먼저 해당되는 애만 수행되고 빠져나간다.
	switch line {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	default:
		fmt.Printf("잘못 입력하셨습니다.")
	}
}
