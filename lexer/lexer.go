package lexer

import (
	"fmt"
	"lc3-dis/bits"
	"lc3-dis/token"
)

type Lexer struct {
	input        []uint16
	instruction  uint16
	line         uint16
	orig         uint16
	position     uint16
	readPosition uint16
	next         func() string
}

func New(input []uint16) *Lexer {
	l := &Lexer{input: input}
	l.readLine()
	return l
}

func (l *Lexer) NextToken() string {
	if l.readPosition == 1 {
		l.readLine()
		return l.GetOrig()
	}
	var str string
	opcodeBits := bits.GetOpcode(l.instruction)

	str = token.BulidInstructionString(opcodeBits, l.instruction)
	l.readLine()
	return str
}

func (l *Lexer) GetOrig() string {
	var str string
	str = fmt.Sprintf(".ORIG x%x", l.orig)
	return str
}

func (l *Lexer) readLine() {
	if int(l.readPosition) >= len(l.input) {
		l.instruction = 0
	} else if int(l.readPosition) == 0 {
		l.instruction = l.input[l.readPosition]
		l.orig = l.instruction
		l.line = l.orig
	} else {
		l.instruction = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

}

func (l *Lexer) string() string {
	return fmt.Sprintf("<Lexer> line=%d, orig=%d, position=%d, readposition=%d, instruction=%d\n> input=%v", l.line, l.orig, l.position, l.readPosition, l.instruction, l.input)

}
