package parser

import (
	"bufio"
	"log"
)

func Function(bufr *bufio.Reader) uint32 {
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}	
	typ := int(b)
	if typ != 0x60 {
		log.Fatalf("Function expected. %x", typ)
	}
	log.Printf("func")
	var c uint32
	c = 0
	c += ValTypeVec(bufr)
	c += ValTypeVec(bufr)
	return c
}
