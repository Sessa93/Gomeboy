package cpu

import (
	"fmt"
)

type Registers struct {
	a Int
	b Int
	c Int
	d Int
	e Int
	h Int
	l Int
	
	sp Int //Stack pointer
	pc Int //Program counter
}

func (Registers r) getA() Int {
	return r.a
}

func (Registers r) getB() Int {
	return r.b
}

func (Registers r) getC() Int {
	return r.c
}

func (Registers r) getD() Int {
	return r.d
}

func (Registers r) getE() Int {
	return r.e
}

func (Registers r) getH() Int {
	return r.h
}

func (Registers r) getL() Int {
	return r.l
}





