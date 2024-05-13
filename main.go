package main

import (
	"bufio"
	"fmt"
	"lc3-dis/lexer"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dir := "/Users/adrianbrady/Workspace/Projects/lc3-dis/"
	fileName := "medium.obj"
	filePath := fmt.Sprintf("%s%s", dir, fileName)

	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	stats, err := file.Stat()
	check(err)

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	if len(bytes)%2 != 0 {
		panic(0)
	}

	var instructions []uint16
	for i := 0; i < len(bytes); i = i + 2 {
		instruction := uint16(bytes[i]) << 8
		instruction = instruction + uint16(bytes[i+1])
		instructions = append(instructions, instruction)
	}

	fmt.Printf("size=%d, bytes=%v\n", size, bytes)
	for _, instr := range instructions {
		fmt.Printf("bin=%b, int=%d\n", instr, instr)
	}

	l := lexer.New(instructions)

	for i, instr := range instructions {
		out := l.NextToken()
		fmt.Printf("lexer[%d] - in=x%x, out=%s\n", i, instr, out)
	}

}
