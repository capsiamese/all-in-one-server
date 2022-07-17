package main

type X struct {
	A int
}

func (X) P() {
	print("x")
}

type W interface {
	P()
}

func testInterface() {
	var n = struct {
		i int
	}{}.i
	var k = W.P
	var m = interface{ P() }.P
	m(X{n})
	k(X{})
	X.P(X{})
}
