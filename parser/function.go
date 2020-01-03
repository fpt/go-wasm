package parser

import (
	"log"
)

func Function(wr *WasmReader) uint32 {
	typ := wr.ReadType()

	if typ != 0x60 {
		log.Fatalf("Function expected. %x", typ)
	}
	log.Printf("func")

	// TODO: Workaround
	bufr := wr.Reader()

	var c uint32
	c = 0
	c += ValTypeVec(bufr)
	c += ValTypeVec(bufr)
	return c
}
