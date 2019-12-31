package parser

import (
	"bufio"
	"log"

	"net.fujlog/go-wasm/valtype"
)

func Import(bufr *bufio.Reader) {
	mod := Name(bufr)
	nm := Name(bufr)
	log.Printf("import mod: %s, name: %s", mod, nm)
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	typ := int(b)
	switch typ {
	case 0x00:
		idx := U32(bufr)
		log.Printf("typeidx %d", idx)
	case 0x01:
		b, err := bufr.ReadByte()
		if err != nil {
			log.Fatalf("Error occured %s", err)
		}
		Limit(bufr)
		log.Printf("tabletype %d", int(b))
	case 0x02:
		Limit(bufr)
		log.Printf("memtype")
	case 0x03:
		log.Printf("global")
		b, err := bufr.ReadByte()
		if err != nil {
			log.Fatalf("Error occured %s", err)
		}
		typ := valtype.ValType(b)
		b, err = bufr.ReadByte()
		if err != nil {
			log.Fatalf("Error occured %s", err)
		}
		mut := int(b)
		log.Printf("globaltype typ:%v mut:%d", typ, mut)
	}
}
