package main

type Character struct {
	power Power
	level Level
	p1    *PointerTest
}

type Power struct {
	name string
}

type Level struct {
	nested NestedNumber
}

type NestedNumber struct {
	number int
}

type PointerTest struct {
	p    *PointerTest
	name string
}

func newPointerTest(n string) *PointerTest {
	var p PointerTest
	p.name = n
	return &p
}
