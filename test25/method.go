package test25

import "fmt"

type Bread struct {
	val string
}

type Jam struct {
}

func (b *Bread) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}

func (b *Bread) String() string {
	return b.val
}

func (j *Jam) GetVal() string {
	return " + jam"
}

func Method() {
	bread := &Bread{val: "bread"}
	jam := &Jam{}

	bread.PutJam(jam)

	fmt.Println(bread)
}
