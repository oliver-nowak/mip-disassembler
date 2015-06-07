package main

import (
	"fmt"
)

var instructions = []uint32{
	0x022DA822, // sub
	0x8EF30018, // lw
	0x12A70004, // beq
	0x02689820,
	0xAD930018,
	0x02697824,
	0xAD8FFFF4,
	0x018C6020,
	0x02A4A825,
	0x158FFFF6,
	0x8E59FFF0}

var func_codes = map[uint32]string{
	0x20: "add",
	0x22: "sub",
	0x24: "and",
	0x25: "or",
	0x2A: "slt"}

var op_codes = map[uint32]string{
	0x04: "beq",
	0x05: "bne",
	0x23: "lw",
	0x2B: "sw"}

var pc uint32 = 0x7A05C

const (
	INSTRUCTION_SIZE uint32 = 0x00000004 // in bytes
	RFORMAT          uint32 = 0x0
	OPCODE_MASK      uint32 = 0xFC000000 // >> 26
	RS_MASK          uint32 = 0x03E00000 // >> 21
	RT_MASK          uint32 = 0x001F0000 // >> 16
	RD_MASK          uint32 = 0x0000F800 // >> 11
	FUNC_MASK        uint32 = 0x0000003F // >> 00
	OFFSET_MASK      uint32 = 0x0000FFFF // >> 00
)

// FORMATS: http://en.wikibooks.org/wiki/MIPS_Assembly/Instruction_Formats#Opcodes

// R-Format:
// op rd, rs, rt

// I-Format
// op rt, rs, offset

func main() {
	fmt.Println("MIPS Disassembler \n")

	showVerbose := false

	for _, instruction := range instructions {

		// increment pc
		pc = pc + 0x4

		// handle opcode format
		if ((instruction & OPCODE_MASK) >> 26) == RFORMAT {
			Do_RFormat(instruction, showVerbose)
		} else {
			Do_IFormat(instruction, showVerbose)
		}
	}
}

func Do_RFormat(instruction uint32, showVerbose bool) {
	opcode := (instruction & OPCODE_MASK) >> 26
	rs := (instruction & RS_MASK) >> 21
	rt := (instruction & RT_MASK) >> 16
	rd := (instruction & RD_MASK) >> 11
	funct := (instruction & FUNC_MASK) >> 0
	func_code := func_codes[funct]

	if showVerbose {
		fmt.Println("---RFORMAT---")
		fmt.Printf("Instruction : 0x%X \n", instruction)
		fmt.Printf("Opcode      : 0x%X \n", opcode)
		fmt.Printf("RS_MASK          : 0x%X \n", rs)
		fmt.Printf("RT_MASK          : 0x%X \n", rt)
		fmt.Printf("RD_MASK          : 0x%X \n", rd)
		fmt.Printf("Func        : 0x%X \n", funct)
		fmt.Println("---END--")
	}

	fmt.Printf("%X     %s  $%d, $%d, $%d \n", pc, func_code, rd, rs, rt)
}

func Do_IFormat(instruction uint32, showVerbose bool) {
	opcode := (instruction & OPCODE_MASK) >> 26
	op := op_codes[opcode]
	rs := (instruction & RS_MASK) >> 21
	rt := (instruction & RT_MASK) >> 16
	offset := (instruction & OFFSET_MASK)
	decompressed_offset := offset << 2

	if showVerbose {
		fmt.Println("---IFORMAT---")
		fmt.Printf("Instruction : 0x%X \n", instruction)
		fmt.Printf("Opcode      : %s \n", op)
		fmt.Printf("RS_MASK          : 0x%X \n", rs)
		fmt.Printf("RT_MASK          : 0x%X \n", rt)
		fmt.Printf("Offset      : 0x%X \n", offset)
		fmt.Printf("D-Offset    : 0x%X \n", decompressed_offset)
		fmt.Println("---END---")
	}

	if op == "lw" || op == "sw" {
		fmt.Printf("%X     %s  $%d, %d ($%d) \n", pc, op, rs, offset, rt)
	} else {
		offset_address := pc + 4 + decompressed_offset
		fmt.Printf("%X     %s  $%d, $%d address $%X \n", pc, op, rt, rs, offset_address)
	}
}
