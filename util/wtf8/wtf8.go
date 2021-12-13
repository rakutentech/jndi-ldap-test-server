package wtf8

import (
	"math"
	"unicode/utf16"
)

func Encode(input string) []byte {
	result := make([]byte, 0, len(input))
	for _, r := range input {
		if r <= math.MaxUint16 {
			result = encodeRune(result, uint16(r))
		} else {
			// Encode surrogate pair as if they were valid unicode codepoints
			surr1, surr2 := utf16.EncodeRune(r)
			result = encodeRune(result, uint16(surr1))
			result = encodeRune(result, uint16(surr2))
		}
	}
	return result
}

const (
	rune1Max = 1<<7 - 1
	rune2Max = 1<<11 - 1

	tx = 0b10000000
	t2 = 0b11000000
	t3 = 0b11100000

	maskx = 0b00111111
)

func encodeRune(p []byte, r uint16) []byte {
	// Negative values are erroneous. Making it unsigned addresses the problem.
	switch {
	case r <= rune1Max:
		return append(p, byte(r))
	case r <= rune2Max:
		p = append(p, t2 | byte(r>>6))
		return append(p, tx | byte(r)&maskx)
	default:
		p = append(p, t3 | byte(r>>12))
		p = append(p, tx | byte(r>>6))
		return append(p, tx | byte(r)&maskx)
	}
}
