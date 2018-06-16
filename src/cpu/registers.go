package cpu

import (
	"fmt"
)

type Registers struct {
	a byte
	b byte
	c byte
	d byte
	e byte
	h byte
	l byte
	
	sp byte //Stack pointer
	pc byte //Program counter
}

func (Registers *r) getA() byte {
	return r.a
}

func (Registers *r) getB() byte {
	return r.b
}

func (Registers *r) getC() byte {
	return r.c
}

func (Registers *r) getD() byte {
	return r.d
}

func (Registers *r) getE() byte {
	return r.e
}

func (Registers *r) getH() byte {
	return r.h
}

func (Registers *r) getL() byte {
	return r.l
}





