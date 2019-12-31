package parser

import (
	"bufio"
	"log"

	"net.fujlog/go-wasm/valtype"
)

func ValTypeVec(bufr *bufio.Reader) uint32 {
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	var n int
	cval := uint32(b)
	vals := make([]byte, cval)
	n, err = bufr.Read(vals)
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	if uint32(n) < cval {
		log.Fatalf("File too short!")
	}
	log.Printf("%d", cval)
	for _, b := range vals {
		valtype.ValType(b)
	}
	return uint32(1) + cval
}
