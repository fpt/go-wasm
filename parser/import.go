package parser

import (
	"log"

	"net.fujlog/go-wasm/valtype"
)

func Import(wr *WasmReader) {
	mod := wr.ReadName()
	nm := wr.ReadName()
	log.Printf("import mod: %s, name: %s", mod, nm)

	typ := wr.ReadType()
	switch typ {
	case 0x00:
		idx := wr.ReadU32()
		log.Printf("typeidx func: %d", idx)
	case 0x01:
		b := wr.ReadByte()
		if b != 0x70 {
			log.Fatalf("Invalid elemtype")
		}

		b = wr.ReadByte()
		wr.ReadLimit()
		log.Printf("tabletype %d", int(b))
	case 0x02:
		wr.ReadLimit()
		log.Printf("memtype")
	case 0x03:
		log.Printf("global")
		b := wr.ReadByte()
		typ := valtype.ValType(b)
		b = wr.ReadByte()
		mut := int(b)
		log.Printf("globaltype typ:%v mut:%d", typ, mut)
	}
}
