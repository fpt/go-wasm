package parser

import (
	"bufio"
	"fmt"
	"log"
)

// u32 LEB128 decoder
func U32(bufr *bufio.Reader) uint32 {
	return uint32(LebUInt(bufr, 32))
}

func LebUInt(bufr *bufio.Reader, nn int) int {
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	n := int(b)
	if n < 0x7F && n < 2^nn {
		return n
	}
	m := LebUInt(bufr, nn-7)
	fmt.Printf(" m = %d\n", m)
	n = 0x7F*m + n - 0x7F
	return n
}