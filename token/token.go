package token

import (
	"bytes"
	"fmt"
	"lc3-dis/bits"
)

const (
	// Opcodes
	ADD    = 0b0001
	AND    = 0b0101
	BR     = 0b0
	JMP    = 0b1100
	JSR    = 0b0100
	LD     = 0b0010
	LDI    = 0b1010
	LDR    = 0b0110
	LEA    = 0b1110
	NOT    = 0b1001
	RTI    = 0b1000
	ST     = 0b0011
	STI    = 0b1011
	STR    = 0b0111
	TRAP   = 0b1111
	UNUSED = 0b1101

	// TRAP Codes
	HALT = 0x25
	IN   = 0x23
	OUT  = 0x21
	GETC = 0x20
	PUTS = 0x22
)

type TokenType uint16

type Token struct {
	TokenType TokenType
	Literal   string
}

func BulidInstructionString(opcode uint8, instruction uint16) string {
	var buffer bytes.Buffer

	switch opcode {
	// ADD  0001 RRR RRR 0 00 RRR
	// ADDi 0001 RRR RRR 1 iiiii
	case ADD:
		buffer.WriteString("ADD ")

		DR := bits.GetDataRegister(instruction)
		SR1 := bits.GetSourceRegister1(instruction)

		buffer.WriteString(fmt.Sprintf("R%d,", DR))
		buffer.WriteString(fmt.Sprintf("R%d,", SR1))
		if bits.GetImm5Bit(instruction) == 0 {
			SR2 := bits.GetSourceRegister2(instruction)
			buffer.WriteString(fmt.Sprintf("R%d", SR2))
		} else {
			imm5 := bits.GetImm5(instruction)
			imm5Str := ConvTwosComp(imm5, 5)
			buffer.WriteString(imm5Str)
		}
		return buffer.String()
	// AND  0101 RRR RRR 0 00 RRR
	// ANDi 0101 RRR RRR 1 iiiii
	case AND:
		buffer.WriteString("AND ")

		DR := bits.GetDataRegister(instruction)
		SR1 := bits.GetSourceRegister1(instruction)

		buffer.WriteString(fmt.Sprintf("R%d,", DR))
		buffer.WriteString(fmt.Sprintf("R%d,", SR1))
		if bits.GetImm5Bit(instruction) == 0 {
			SR2 := bits.GetSourceRegister2(instruction)
			buffer.WriteString(fmt.Sprintf("R%d", SR2))
		} else {
			imm5 := bits.GetImm5(instruction)
			imm5Str := ConvTwosComp(imm5, 5)
			buffer.WriteString(imm5Str)
		}
		return buffer.String()
	// BR 0000 nzp x xxxx xxxx
	case BR:
		nzp := bits.GetNZP(instruction)
		offset9 := bits.GetPCOffset9(instruction)
		offset9Str := ConvTwosComp(uint16(offset9), 9)

		var n string
		var z string
		var p string

		if nzp[0] == 0 {
			n = ""
		} else {
			n = "n"
		}

		if nzp[1] == 0 {
			z = ""
		} else {
			z = "z"
		}

		if nzp[2] == 0 {
			p = ""
		} else {
			p = "p"
		}

		buffer.WriteString("BR")
		buffer.WriteString(fmt.Sprintf("%s%s%s ", n, z, p))
		buffer.WriteString(offset9Str)

		return buffer.String()
	// JMP 1100 000 RRR 000000
	// RET 1100 000 111 000000
	case JMP:
		baseR := int(bits.GetSourceRegister1(instruction))

		if baseR == 0b111 {
			buffer.WriteString("RET")
		} else {
			buffer.WriteString(fmt.Sprintf("JMP R%d", baseR))
		}

		return buffer.String()
	// JSR  0100 1 xxx xxxx xxxx
	// JSRR 0100 0 00 RRR 000000
	case JSR:
		n := bits.GetNZP(instruction)[0]
		baseR := bits.GetSourceRegister1(instruction)
		offset11 := bits.GetPCOffset11(instruction)

		if n == 0 {
			buffer.WriteString(fmt.Sprintf("JSRR R%d", baseR))
		} else {
			buffer.WriteString("JSR ")
			offset11Str := ConvTwosComp(uint16(offset11), 11)
			buffer.WriteString(offset11Str)
		}

		return buffer.String()
	// LD 0010 RRR x xxxx xxxx
	case LD:
		register := bits.GetDataRegister(instruction)
		offset9 := bits.GetPCOffset9(instruction)
		offset9Str := ConvTwosComp(uint16(offset9), 9)

		buffer.WriteString(fmt.Sprintf("LD "))
		buffer.WriteString(fmt.Sprintf("R%d,", register))
		buffer.WriteString(offset9Str)

		return buffer.String()
	// LDI 1010 RRR x xxxx xxxx
	case LDI:
		register := bits.GetDataRegister(instruction)
		offset9 := bits.GetPCOffset9(instruction)
		offset9Str := ConvTwosComp(uint16(offset9), 9)

		buffer.WriteString(fmt.Sprintf("LDI "))
		buffer.WriteString(fmt.Sprintf("R%d,", register))
		buffer.WriteString(offset9Str)

		return buffer.String()
	// LDR 0110 RRR RRR xxxxxx
	case LDR:
		dataRegister := bits.GetDataRegister(instruction)
		baseRegister := bits.GetSourceRegister1(instruction)
		offset6 := bits.GetOffset6(instruction)
		offset6Str := ConvTwosComp(uint16(offset6), 6)

		buffer.WriteString(fmt.Sprintf("LDR "))
		buffer.WriteString(fmt.Sprintf("R%d,", dataRegister))
		buffer.WriteString(fmt.Sprintf("R%d,", baseRegister))
		buffer.WriteString(offset6Str)

		return buffer.String()

	// LEA 1110 RRR x xxxx xxxx
	case LEA:
		register := bits.GetDataRegister(instruction)
		offset9 := bits.GetPCOffset9(instruction)
		offset9Str := ConvTwosComp(uint16(offset9), 9)

		buffer.WriteString(fmt.Sprintf("LEA "))
		buffer.WriteString(fmt.Sprintf("R%d,", register))
		buffer.WriteString(offset9Str)

		return buffer.String()
	// NOT 1001 RRR RRR 1 11111
	case NOT:
		dataRegister := bits.GetDataRegister(instruction)
		baseRegister := bits.GetSourceRegister1(instruction)
		offset6 := bits.GetOffset6(instruction)

		if offset6 != 0b111111 {
			return "ILLEGAL"
		}

		buffer.WriteString(fmt.Sprintf("NOT "))
		buffer.WriteString(fmt.Sprintf("R%d,", dataRegister))
		buffer.WriteString(fmt.Sprintf("R%d", baseRegister))

		return buffer.String()
	// RTI 1000 0000 0000 0000
	case RTI:
		if instruction != 0b1000_0000_0000_0000 {
			return "ILLEGAL"
		} else {
			return "RTI"
		}
	// ST 0011 RRR x xxxx xxxx
	case ST:
		register := bits.GetDataRegister(instruction)
		offset9 := bits.GetPCOffset9(instruction)
		offset9Str := ConvTwosComp(uint16(offset9), 9)

		buffer.WriteString(fmt.Sprintf("ST "))
		buffer.WriteString(fmt.Sprintf("R%d,", register))
		buffer.WriteString(offset9Str)

		return buffer.String()
	// STI 1011 RRR x xxxx xxxx
	case STI:
		register := bits.GetDataRegister(instruction)
		offset9 := bits.GetPCOffset9(instruction)
		offset9Str := ConvTwosComp(uint16(offset9), 9)

		buffer.WriteString(fmt.Sprintf("STI "))
		buffer.WriteString(fmt.Sprintf("R%d,", register))
		buffer.WriteString(offset9Str)

		return buffer.String()
	// STR 0111 RRR RRR xxxxxx
	case STR:
		dataRegister := bits.GetDataRegister(instruction)
		baseRegister := bits.GetSourceRegister1(instruction)
		offset6 := bits.GetOffset6(instruction)
		offset6Str := ConvTwosComp(uint16(offset6), 6)

		buffer.WriteString(fmt.Sprintf("STR "))
		buffer.WriteString(fmt.Sprintf("R%d,", dataRegister))
		buffer.WriteString(fmt.Sprintf("R%d,", baseRegister))
		buffer.WriteString(offset6Str)

		return buffer.String()
	// TRAP 1111 0000 xxxx xxxx
	case TRAP:
		top8 := bits.GetRange(0, 7, instruction)

		if top8 != 0b1111_0000 {
			return "ILLEGAL"
		}

		trapvect8 := bits.GetTrapVect8(instruction)

		switch trapvect8 {
		case HALT:
			buffer.WriteString("HALT")
		case IN:
			buffer.WriteString("IN")
		case OUT:
			buffer.WriteString("OUT")
		case GETC:
			buffer.WriteString("GETC")
		case PUTS:
			buffer.WriteString("PUTS")
		default:
			buffer.WriteString(fmt.Sprintf("TRAP x%x", trapvect8))
		}

		return buffer.String()
	// UNUSED 1101 xxxx xxxx xxxx
	case UNUSED:
		return "UNUSED"
	default:
		return "ILLEGAL"
	}
}

func ConvTwosComp(number uint16, bitLength uint8) string {
	var twos int
	shiftLength := bitLength - 1
	negBit := number >> shiftLength
	if negBit == 1 {
		negative := (number >> uint16(shiftLength)) << (uint16(shiftLength) + 1)
		number = (number << (15 - bitLength)) >> (15 - bitLength)
		twos = -int(negative - number)
		return fmt.Sprintf("#%d", twos)

	} else {
		return fmt.Sprintf("#%d", number)
	}
}
