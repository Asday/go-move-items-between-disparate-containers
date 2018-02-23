package main

import "github.com/davecgh/go-spew/spew"

type base []int

type child1 base
type child2 base

func moveItemsBase(from, to base, amount int) (base, base) {
	for _, item := range from[:3] {
		to = append(to, item)
	}

	spew.Dump(from, to)

	return from, to
}

func main() {
	c1 := child1{1, 2, 3, 4, 5}
	c2 := child2{}

	// There must be a better way!
	out1, out2 := moveItemsBase(base(c1), base(c2), 3)
	c1 = child1(out1)
	c2 = child2(out2)

	spew.Dump(c1, c2)
}
