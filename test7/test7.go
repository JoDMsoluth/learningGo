package test7

import "fmt"

func Test7() {
	// 배열의 길이는 Type크기 * 개수
	// String은 1byte배열로 나타낼 수 있지만
	// rune의 배열로 나타낼 수도 있다(UTF-8의 배열)
	// 다른 언어는 ASCII 코드로 다루기 떄문에 어렵지는 않다.
	s := "Hello 월드"

	s2 := []rune(s)

	fmt.Println("len(s2) = ", len(s2))

	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i], " ", string(s2[i]))
	}

	// 배열의 복사
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < len(arr); i++ {
		clone[i] = arr[i]
	}

	fmt.Println(clone)

	// 배열의 역순
	for i := 0; i < len(arr); i++ {
		clone[i] = arr[len(arr)-i-1]
	}
	fmt.Println(clone)

	// 빠른 배열의 역순, (반절만 돌고 복사하지도 않음)
	for i := 0; i < len(arr)/2; i++ {
		//
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println(arr)

	// radix sort
	// 특정한 경우에만 사용
	// 각 원소의 범위가 정수여야하며 작아야하고 요소가 정해져야된다.
	// 예 ) 0 ~ 9의 정수만 들어간다.
	arr2 := [10]int{0, 2, 5, 7, 2, 3, 6, 4, 1, 8}
	temp := [10]int{}

	// temp에 각 원소들이 몇변 나오는지 저장한다.
	for i := 0; i < len(arr2); i++ {
		idx := arr2[i]
		temp[idx]++
	}
	fmt.Println(temp)

	idx := 0
	// 0이 두개면 2번 반복해야 된다 때문에, 2중반복문
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr2[idx] = i
			idx++
		}
	}
	fmt.Println(arr2)
}
