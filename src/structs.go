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

	var deepCopyCandidate Character
	deepCopyCandidate = c
	deepCopyCandidate.power.name = "Shallow Coding Skills"
	deepCopyCandidate.level.nested.number = 2
	deepCopyCandidate.p1.p.name = "shallow pointer"
	fmt.Println("original object after shallow copy:\n\n" + render.Render(c) + "\n")
	fmt.Println("shallow copy:\n\n" + render.Render(deepCopyCandidate) + "\n")

	var shallowCopy *Character
	shallowCopy = &c
	shallowCopy.power.name = "Mad Coding Skills"
	shallowCopy.level.nested.number = 3
	shallowCopy.p1.name = "deep pointer"
	fmt.Println("original object after deep copy:\n\n" + render.Render(c) + "\n")
	fmt.Println("deep copy:\n\n" + render.Render(shallowCopy) + "\n")
	fmt.Println("conclusion:\n\nit can be deceitful to assign variables when trying to deep copy, as most " +
		"objects can be changed without affecting the original memory addresses, but when copying a pointer it will " +
		"still keep pointing to the same memory address")
}
