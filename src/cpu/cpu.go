package cpu

import {
	"registers"
	"opcodes"
}

type Cpu struct {
	registers Registers
	opcodes Opcodes
	speed int
	interruptsEnabled bool
}

func NewCPU() {
	
}

func (Cpu *cpu) Reset() {
	cpu.registers.A = 0
	cpu.registers.B = 0
	cpu.registers.C = 0
	cpu.registers.D = 0
	cpu.registers.E = 0
	cpu.registers.F = 0
	cpu.registers.H = 0
	cpu.registers.L = 0

	cpu.registers.PC = 0
	cpu.registers.SP = 0

	cpu.speed = 1
}

func (Cpu *) Step() {

}