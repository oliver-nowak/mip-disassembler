package main

import (
	"fmt"
)

var instructions = []uint32{
	0x022DA822}

// 0x8EF30018,
// 0x12A70004,
// 0x02689820,
// 0xAD930018,
// 0x02697824,
// 0xAD8FFFF4,
// 0x018C6020,
// 0x02A4A825,
// 0x158FFFF6,
// 0x8E59FFF0}

const (
	START_ADDRESS uint32 = 0x7A060
	OPCODE        uint32 = 0xFC000000 // >> 26
	REG_1         uint32 = 0x03E00000 // >> 21
	REG_2         uint32 = 0x001F0000 // >> 16
	REG_3         uint32 = 0x0000F800 // >> 11
	SHAMT         uint32 = 0x000007C0 // >> 06
	FUNC          uint32 = 0x0000003F // >> 00
	OFFSET        uint32 = 0x0000FFFF // >> 00
)

func main() {
	fmt.Println("hello, world\n ", len(instructions))

	for index, instruction := range instructions {
		fmt.Printf("Instruction #%d : 0x%X \n", index, instruction)

		opcode := (instruction & OPCODE) >> 26
		r1 := (instruction & REG_1) >> 21
		r2 := (instruction & REG_2) >> 16
		r3 := (instruction & REG_3) >> 11
		shamt := (instruction & SHAMT) >> 6
		funct := (instruction & FUNC) >> 0

		fmt.Printf("Opcode : 0x%X \n", opcode)
		fmt.Printf("R1     : 0x%X \n", r1)
		fmt.Printf("R2     : 0x%X \n", r2)
		fmt.Printf("R3     : 0x%X \n", r3)
		fmt.Printf("SHAMT  : 0x%X \n", shamt)
		fmt.Printf("Func   : 0x%X \n", funct)
	}
}
