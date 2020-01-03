package parser

import (
	"log"

	"net.fujlog/go-wasm/valtype"
)

func Import(wr *WasmReader) {

	// TODO: Workaround
	bufr := wr.Reader()

	mod := wr.ReadName()
	nm := wr.ReadName()
	log.Printf("import mod: %s, name: %s", mod, nm)

	typ := wr.ReadType()
	switch typ {
	case 0x00:
		idx := wr.ReadU32()
		log.Printf("typeidx func: %d", idx)
	case 0x01:
		b, err := bufr.ReadByte()
		if err != nil {
			log.Fatalf("Error occured %s", err)
		}
		wr.ReadLimit()
		log.Printf("tabletype %d", int(b))
	case 0x02:
		wr.ReadLimit()
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
