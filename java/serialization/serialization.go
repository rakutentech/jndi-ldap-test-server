package javaser

import (
	"encoding/binary"
	"github.com/rakuten-tech/jndi-ldap-test-server/util/wtf8"
	"math"
)

const headerSize = 5

func EncodeString(input string) []byte {
	encodedStr := wtf8.Encode(input)
	encodedLen := len(encodedStr)
	isLong := encodedLen > math.MaxUint16
	var lengthSize int
	var typeCode byte
	if isLong {
		lengthSize = 8
		typeCode = 0x75 // Long UTF-8 String
	} else {
		lengthSize = 2
		typeCode = 0x74 // Short UTF-8 String
	}

	result := make([]byte, headerSize+lengthSize+encodedLen)
	result[0] = 0xac // MAGIC part 1
	result[1] = 0xed // MAGIC part 2
	result[2] = 0x00 // Version part 1
	result[3] = 0x05 // Version part 2
	result[4] = typeCode

	b := result[headerSize:]
	if isLong {
		binary.BigEndian.PutUint64(b, uint64(encodedLen))
		b = b[8:]
	} else {
		binary.BigEndian.PutUint16(b, uint16(encodedLen))
		b = b[2:]
	}

	copy(b, encodedStr)

	return result
}
