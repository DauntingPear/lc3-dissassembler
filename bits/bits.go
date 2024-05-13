package bits

func GetOpcode(instruction uint16) uint8 {
	bits := (instruction >> uint16(12))
	return uint8(bits)
}

func GetDataRegister(instruction uint16) uint8 {
	bits := (instruction << 4)
	bits = (bits >> 13)
	return uint8(bits)
}

func GetSourceRegister1(instruction uint16) uint8 {
	bits := (instruction << 7)
	bits = (bits >> 13)
	return uint8(bits)
}

func GetSourceRegister2(instruction uint16) uint8 {
	bits := (instruction << 13)
	bits = (bits >> 13)
	return uint8(bits)
}

func GetImm5(instruction uint16) uint16 {
	bits := (instruction << 11)
	bits = (bits >> 11)
	return bits
}

func GetImm5Bit(instruction uint16) uint8 {
	bits := (instruction << 10)
	bits = (bits >> 15)
	return uint8(bits)
}

func GetNZP(instruction uint16) [3]uint8 {
	bits := (instruction << 4)
	bits = (bits >> 13)
	n := uint8((bits & 0b100) >> 2)
	z := uint8((bits & 0b010) >> 1)
	p := uint8(bits & 0b001)
	nzp := [3]uint8{n, z, p}
	return nzp
}

func GetPCOffset9(instruction uint16) uint8 {
	bits := (instruction << 7)
	bits = (bits >> 7)
	return uint8(bits)
}

func GetPCOffset11(instruction uint16) uint8 {
	bits := (instruction << 5)
	bits = (bits >> 5)
	return uint8(bits)
}

func GetOffset6(instruction uint16) uint8 {
	bits := (instruction << 10)
	bits = (bits >> 10)
	return uint8(bits)
}

func GetTrapVect8(instruction uint16) uint8 {
	bits := (instruction << 8)
	bits = (bits >> 8)
	return uint8(bits)
}

// get 2,5
// 0123_4567
// 0011_0100
// 1101_0000
// 0000_1101

func GetRange(start int, end int, instruction uint16) uint16 {
	if start >= end || start > 15 || end > 15 || start < 0 || end < 0 {
		return 0
	}
	leftShift := start
	rightShift := leftShift + (15 - end)
	bits := (instruction << leftShift)
	bits = (bits >> rightShift)

	return bits

}
