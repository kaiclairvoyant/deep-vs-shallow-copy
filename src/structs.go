package main

import (
	"fmt"
	"github.com/luci/go-render/render"
)

func copyStructs() {
	fmt.Println("____________Initiating struct copies____________")
	nested := NestedNumber{1}
	l := Level{nested}
	n := Power{"Super Intelligence"}
	p := newPointerTest("first pointer")
	p.p = newPointerTest("second pointer")
	c := Character{n, l, p}

	fmt.Println("original object:\n" + render.Render(c) + "\n")

	var testCopy Character
	testCopy = c
	testCopy.power.name = "Shallow Copy"
	testCopy.level.nested.number = 2
	testCopy.p1.p.name = "shallow pointer"
	fmt.Println("original object after shallow copy:\n" + render.Render(c) + "\n")
	fmt.Println("shallow copy:\n" + render.Render(testCopy) + "\n")

	fmt.Println("conclusion:\nit can be deceitful to assign variables when trying to create a deep copy, as most " +
		"structs can be changed without affecting the original memory addresses, but when copying a pointer it will " +
		"still keep pointing to the same memory address of the original pointer")
}
