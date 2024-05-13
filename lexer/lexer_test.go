package lexer

import (
	"testing"
)

func TestOpcodes(t *testing.T) {
	input := []uint16{
		0b0011_0000_0000_0000,  // .ORIG x3000
		0b0001_100_010_000_100, // ADD R4,R2,R4
		0b0001_100_010_1_00100, // ADD R4,R2,#4
		0b0001_100_010_1_11100, // ADD R4,R2,#-4
		0b0101_111_101_000_001, // AND R7,R5,R1
		0b0101_110_011_1_00111, // AND R6,R3,#7
		0b0000_111_0_1100_0000, // BRnzp #192
		0b1100_000_010_000000,  // JMP R2
		0b1100_000_111_000000,  // RET
		0b0100_1_000_0000_0110, // JSR #6
		0b0100_000_010_000000,  // JSRR R2
		0b0010_000_0_0000_0011, // LD R0,#3
		0b1010_001_0_0000_0111, // LDI R1,#7
		0b0110_011_001_000100,  // LDR R3,R1,#4
		0b1110_010_0_0000_1100, // LEA R2,#12
		0b1000_0000_0000_0000,  // RTI
		0b0011_000_0_0000_0010, // ST R0,#2
		0b1011_010_0_0000_1100, // STI R2,#12
		0b0111_011_111_001100,  // STR R3,R7,#12
		0b1111_0000_0010_0101,  // HALT
		0b1111_0000_0010_0011,  // IN
		0b1111_0000_0010_0001,  // OUT
		0b1111_0000_0010_0000,  // GETC
		0b1111_0000_0010_0010,  // PUTS
		0b1111_0000_0000_0001,  // TRAP 0x1
		0b1101_0000_0000_0000,  // UNUSED
		0b1111_1000_0000_0000,  // ILLEGAL
		0b0001_100_010_1_11111, // ADD R4,R2,#-1
	}

	tests := []string{
		".ORIG x3000",
		"ADD R4,R2,R4",
		"ADD R4,R2,#4",
		"ADD R4,R2,#-4",
		"AND R7,R5,R1",
		"AND R6,R3,#7",
		"BRnzp #192",
		"JMP R2",
		"RET",
		"JSR #6",
		"JSRR R2",
		"LD R0,#3",
		"LDI R1,#7",
		"LDR R3,R1,#4",
		"LEA R2,#12",
		"RTI",
		"ST R0,#2",
		"STI R2,#12",
		"STR R3,R7,#12",
		"HALT",
		"IN",
		"OUT",
		"GETC",
		"PUTS",
		"TRAP x1",
		"UNUSED",
		"ILLEGAL",
		"ADD R4,R2,#-1",
	}

	l := New(input)

	for i, tt := range tests {
		out := l.NextToken()
		if out != tt {
			t.Errorf("tests[%d] - string wrong. expected=%s, got=%s", i, tt, out)
		}
	}

}
