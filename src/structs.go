package main

import (
	"fmt"
	"github.com/luci/go-render/render"
)

func copyStructs() {
	nested := NestedNumber{1}
	l := Level{nested}
	n := Power{"Super Intelligence"}
	p := newPointerTest("first pointer")
	p.p = newPointerTest("second pointer")
	c := Character{n, l, p}

	fmt.Println("original object:\n\n" + render.Render(c) + "\n")

	var fakeDeep Character
	fakeDeep = c
	fakeDeep.power.name = "Fake Deep Coding Skills"
	fakeDeep.level.nested.number = 2
	fakeDeep.p1.p.name = "fake deep pointer"
	fmt.Println("original object after fake deep copy:\n\n" + render.Render(c) + "\n")
	fmt.Println("fake deep copy:\n\n" + render.Render(fakeDeep) + "\n")

	var shallowCopy *Character
	shallowCopy = &c
	shallowCopy.power.name = "Mad Coding Skills"
	shallowCopy.level.nested.number = 3
	shallowCopy.p1.name = "shallow pointer"
	fmt.Println("original object after shallow copy:\n\n" + render.Render(c) + "\n")
	fmt.Println("shallow copy:\n\n" + render.Render(shallowCopy) + "\n")
	fmt.Println("conclusion:\n\nit can be deceitful to assign variables when trying to shallow copy, as most " +
		"objects can be changed without affecting the original memory addresses, but when copying a pointer it will " +
		"still keep pointing to the same memory address")
}
