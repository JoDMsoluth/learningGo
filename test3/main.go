package main

import "fmt"

func main() {
	a := 4
	b := 2

	// a&b : 0100 AND 0010 = 0000
	// a|b : 0100 OR 0010 = 0110
	// a^b : 0100 XOR 0010 = 0110
	fmt.Printf("a&b : %v\n", a&b)
	fmt.Printf("a|b : %v\n", a|b)
	fmt.Printf("a^b : %v\n", a^b)

	// Shift
	// a>>4 : 00000100 >> 4 = 00000000
	// a<<2 : 00000100 << 2 = 00010000
	fmt.Printf("a>>4 : %v\n", a>>4)
	fmt.Printf("a<<2 : %v\n", a<<2)

	c := false
	if c {
		fmt.Println("참")
	}
	fmt.Println("거짓")
}
