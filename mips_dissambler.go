package main

import (
	"fmt"
)

var instructions = []uint32{
	0x022DA822, // add
	0x8EF30018, // lw
	0x12A70004} // beq
// 0x02689820,
// 0xAD930018,
// 0x02697824,
// 0xAD8FFFF4,
// 0x018C6020,
// 0x02A4A825,
// 0x158FFFF6,
// 0x8E59FFF0}

var func_codes = map[uint32]string{
	0x20: "add",
	0x22: "sub",
	0x24: "and",
	0x25: "or",
	0x2A: "slt"}

var op_codes = map[uint32]string{
	0x23: "lw",
	0x2B: "sw",
	0x04: "beq",
	0x05: "bne"}

var pc = START_ADDRESS

const (
	START_ADDRESS    uint32 = 0x7A060
	INSTRUCTION_SIZE uint32 = 0x4
	OPCODE           uint32 = 0xFC000000 // >> 26
	RS               uint32 = 0x03E00000 // >> 21
	RT               uint32 = 0x001F0000 // >> 16
	RD               uint32 = 0x0000F800 // >> 11
	SHAMT            uint32 = 0x000007C0 // >> 06
	FUNC             uint32 = 0x0000003F // >> 00
	OFFSET           uint32 = 0x0000FFFF // >> 00
	RFORMAT          uint32 = 0x0
)

// R-Format
// op rd, rs, rt

// I-Format
// op rs, rt, offset

func main() {
	fmt.Println("MIPS Disassembler \n")

	for index, instruction := range instructions {
		fmt.Printf("Instruction #%d : 0x%X \n", index, instruction)

		opcode := (instruction & OPCODE) >> 26

		if opcode == RFORMAT {
			fmt.Println("---RFORMAT---")

			rs := (instruction & RS) >> 21
			rt := (instruction & RT) >> 16
			rd := (instruction & RD) >> 11
			shamt := (instruction & SHAMT) >> 6
			funct := (instruction & FUNC) >> 0
			func_code := func_codes[funct]

			fmt.Printf("Opcode : 0x%X \n", opcode)
			fmt.Printf("RS     : 0x%X \n", rs)
			fmt.Printf("RT     : 0x%X \n", rt)
			fmt.Printf("RD     : 0x%X \n", rd)
			fmt.Printf("SHAMT  : 0x%X \n", shamt)
			fmt.Printf("Func   : 0x%X \n", funct)
			fmt.Println("---END--")

			fmt.Printf("%X     %s  $%d, $%d, $%d \n", pc, func_code, rd, rs, rt)
		} else {
			fmt.Println("---IFORMAT---")

			op := op_codes[opcode]
			rs := (instruction & RS) >> 21
			rt := (instruction & RT) >> 16
			offset := (instruction & OFFSET)
			decompressed_offset := offset << 2

			fmt.Printf("Opcode  : %s \n", op)
			fmt.Printf("RS      : 0x%X \n", rs)
			fmt.Printf("RT      : 0x%X \n", rt)
			fmt.Printf("Offset  : 0x%X \n", offset)
			fmt.Printf("D-Offset: 0x%X \n", decompressed_offset)
			fmt.Println("---END---")

			// increment pc
			pc += INSTRUCTION_SIZE

			if op == "lw" || op == "sw" {
				fmt.Printf("%X     %s  $%d, %d ($%d) \n", pc, op, rs, offset, rt)
			} else {
				offset_address := pc + INSTRUCTION_SIZE + decompressed_offset
				fmt.Printf("%X     %s  $%d, $%d address $%X \n", pc, op, rs, rt, offset_address)
			}

		}
	}
}
