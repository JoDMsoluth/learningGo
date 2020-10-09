package test27

import (
	"fmt"
	"strconv"
)

// 오리처럼 말하고 오리처럼 행동하면 그건 오리다.
// duck type : go, python

type StructA struct {
}

func (a *StructA) AAA(x int) int {
	return x * x
}

func (a *StructA) BBB(x int) string {
	return "X=" + strconv.Itoa(x)
}

type StructB struct {
}

func (b *StructB) AAA(x int) int {
	return x * x
}

func (b *StructB) BBB(x int) string {
	return "X=" + strconv.Itoa(x)
}

func Test27() {
	var c InterfaceA
	c = &StructA{}

	var b InterfaceA
	b = &StructB{}

	fmt.Println(c.BBB(3))
	fmt.Println(b.AAA(3))
}
