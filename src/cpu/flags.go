package cpu

import (
	"fmt"
)

type Flags struct {
	Z bool // Zero bit
	N bool // Subtract bit
	H bool // Half Carry bit
	C bool // Carry bit
}

func (Flags *f) isZ() bool {
	return f.Z
}

func (Flags *f) isN() bool {
	return f.N
}

func (Flags *f) isH() bool {
	return f.H
}

func (Flags *f) isC() bool {
	return f.C
}


func (Flags *f) setZ() {
	f.Z := true
}

func (Flags *f) setN() {
	f.N := true
}

func (Flags *f) setH() {
	f.H := true
}

func (Flags *f) setC() {
	f.C := true
}


func (Flags *f) clearZ() {
	f.Z := false
}

func (Flags *f) clearN() {
	f.N := false
}

func (Flags *f) clearH() {
	f.H := false
}

func (Flags *f) clearC() {
	f.C := false
}