package parser

import (
	"log"
)

func Section(wr *WasmReader) {
	typ := wr.ReadType()
	secsize := wr.ReadU32()

	log.Printf("section size: %d", secsize)
	if typ == 1 {
		veclen := wr.ReadU32()
		log.Printf("funcvec len: %d", veclen)
		for veclen > 0 {
			Function(wr)
			veclen -= 1
		}
	} else if typ == 2 {
		veclen := wr.ReadU32()
		log.Printf("importvec len: %d", veclen)
		for veclen > 0 {
			Import(wr)
			veclen -= 1
		}
	} else if typ == 3 {
		veclen := wr.ReadU32()
		log.Printf("funcsec len: %d", veclen)
		for veclen > 0 {
			typeidx := wr.ReadU32()
			log.Printf("typeidx: %d", typeidx)
			veclen -= 1
		}
	} else if typ == 4 {
		log.Printf("tablesec")
	} else if typ == 5 {
		log.Printf("memorysec")
	} else if typ == 6 {
		log.Printf("globalsec")
	} else if typ == 7 {
		veclen := wr.ReadU32()
		log.Printf("exportsec veclen: %d", veclen)
		for veclen > 0 {
			nm := wr.ReadName()
			log.Printf("name: %s", nm)
			idx := wr.ReadU32()
			log.Printf("idx: %d", idx)
			if idx == 0 {
				funcidx := wr.ReadU32()
				log.Printf("funcidx: %d", funcidx)
			}
			veclen -= 1
		}
	} else if typ == 8 {
		log.Printf("startsec")
	} else if typ == 9 {
		log.Printf("elementsec")
	} else if typ == 10 {
		veclen := wr.ReadU32()
		log.Printf("codesec veclen: %d", veclen)
		for veclen > 0 {
			size := wr.ReadU32()
			log.Printf("codesec size: %d", size)
			nlocals := wr.ReadU32()
			log.Printf("codesec nlocals: %d", nlocals)
			for nlocals > 0 {
				n := wr.ReadU32()
				log.Printf("codesec n: %d", n)
				t := wr.ReadU32()
				log.Printf("codesec t: %d", t)
				nlocals -= 1
			}
			Expr(wr)
			veclen -= 1
		}
	} else if typ == 11 {
		veclen := wr.ReadU32()
		log.Printf("datasec len: %d", veclen)
		for veclen > 0 {
			memidx := wr.ReadU32()
			log.Printf("datasec memidx: %d", memidx)
			Expr(wr)
			vecb := wr.ReadU32()
			log.Printf("datasec vecb: %d", vecb)

			// TODO: workaround
			bufr := wr.Reader()
			for vecb > 0 {
				b, err := bufr.ReadByte()
				if err != nil {
					log.Fatalf("Error occured %s", err)
				}
				log.Printf("byte: %x", b)
				vecb -= 1
			}
			veclen -= 1
		}
	}
}
