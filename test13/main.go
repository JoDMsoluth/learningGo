package main

import "fmt"

// slice는 하나의 struct로 볼 수 있다.
// 구성요소는
// {
//	 pointer * (시작주소)
//	 len 	int
//	 cap 	int
// }
//

func main() {
	var s []int
	fmt.Printf("S주소 : %p\n", s)

	s = make([]int, 3)
	fmt.Printf("make S주소 %p\n", s)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	// cap이 부족해서 새로운 배열을 생성!!
	// 거기에 카피하고 반환한다.
	// t는 새로 복사한 배열의 주소를 가진다.
	t := append(s, 400, 500, 600, 700)
	fmt.Printf("T주소 : %p\n", t)

	fmt.Println("S: ", s, len(s), cap(s))
	fmt.Println("T: ", t, len(t), cap(t))

	// cap이 부족하지 않기 때문에 그냥 t를 바로 사용한다.
	u := append(t, 800)
	fmt.Println("U: ", u, len(u), cap(u))

	u[1] = 0
	u[3] = 0
	u[5] = 0
	u[7] = 0
	fmt.Println("/////////////////////")
	fmt.Println("S: ", s, len(s), cap(s))
	fmt.Println("T: ", t, len(t), cap(t))
	fmt.Println("U: ", u, len(u), cap(u))
	fmt.Printf("S주소 : %p\n", s)
	fmt.Printf("T주소 : %p\n", t)
	fmt.Printf("U주소 : %p\n", u)
}
