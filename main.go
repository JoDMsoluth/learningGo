package main

import (
	"fmt"
	"strings"

	"github.com/hyehyeong/learnGo/something"
)

// 패키지 함수가 대문자인 이유
// Go에서는 function을 export하기 위해서는 대문자로 해야 하기 때문

func multiply(a, b int) int {
	return a * b
}

// ... 을 써서 여러개의 string argment를 받을 수 있다.
func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Go에서 함수는 여러 개를 반환할 수 있다.
// naked return : 구지 작성하지 않아도 알아서 리턴해준다.
func lenAndUpper(name string) (length int, uppercase string) {
	// defer는 함수가 끝나고 실행된다. - api 호출 후 후속처리, 이미지를 열거나, 파일을 생성한 후 등...
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// if문 사용법
func canIDrink(age int) bool {
	// variable expression : if문 안에 변수를 만들어 사용할 수 있다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

// switch if-else를 난무하는 걸 피하고 싶을 때
func switchCanIDrink(age int) bool {
	// switch의 variable expression
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

// pointer
func pointer() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a, &a, b, *b)
}

func array() {
	// [3] 혹은 []로 지정할 수 잇고 암시적으로 지정할 수 있다.
	names := []string{"nico", "lynn", "dal"}
	// names[3] = "alala"
	// names[4] = "bddddd"

	// [] 처럼 지정한 걸 slice라고 한다.
	// slice로 지정했을 때 추가 요소를 넣고 싶으면 append() 함수를 사용
	// concat같은 거
	names = append(names, "flynn")

	fmt.Println(names)
}

func mapExample() {
	// map : key-value가 있고 js의 object와 비슷
	nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}
	fmt.Println(nico)
}

// map과 비슷하지만 value가 다양한 타입을 쓰고 싶을 때
type person struct {
	name    string
	age     int
	favFood []string
}

func structExample() {
	favFood := []string{"apple", "banana"}

	// 안티패턴 -> nico := person{"nico", 18, favFood}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.name, nico.age, nico.favFood)
}

func main() {
	fmt.Println("Hello World")
	something.SayHello()

	// 대문자가 아니라 에러가 남
	// something.sayBye()

	// Go가 알아서 타입을 정해줌
	// 축약형은 function 안에서만 가능
	age := 12

	var name bool = false
	name = true

	fmt.Println(name, age)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)

	repeatMe("nico", "lynn", "da", "marl", "flynn")

	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	fmt.Println(canIDrink(16))

	a := 2
	b := a
	a = 10
	fmt.Println(a, b)

	pointer()

	array()

	mapExample()

	structExample()
}
