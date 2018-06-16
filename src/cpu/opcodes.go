package cpu

import {
	"registers"
}

type Opcode {
	Signature       byte
	Description  string
	OperandsSize int
	Cycles       int
	Operands     [2]byte
}

var mapping = map[byte]Opcode {
	0x83: 
}

function LDnnn()

