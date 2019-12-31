package parser

import (
	"bufio"
	"log"
)

func Section(bufr *bufio.Reader) {
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	typ := int(b)
	secsize := U32(bufr)
	log.Printf("section size: %d", secsize)
	if typ == 1 {
		veclen := U32(bufr)
		log.Printf("funcvec len: %d", veclen)
		for secsize > 0 && veclen > 0 {
			secsize -= Function(bufr)
			veclen -= 1
		}
	} else if typ == 2 {
		veclen := U32(bufr)
		log.Printf("importvec len: %d", veclen)
		for veclen > 0 {
			Import(bufr)
			veclen -= 1
		}
	} else if typ == 3 {
		veclen := U32(bufr)
		log.Printf("funcsec len: %d", veclen)
		for veclen > 0 {
			typeidx := U32(bufr)
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
		veclen := U32(bufr)
		log.Printf("exportsec veclen: %d", veclen)
		for veclen > 0 {
			nm := Name(bufr)
			log.Printf("name: %s", nm)
			idx := U32(bufr)
			log.Printf("idx: %d", idx)
			if idx == 0 {
				funcidx := U32(bufr)
				log.Printf("funcidx: %d", funcidx)
			}
			veclen -= 1
		}
	} else if typ == 8 {
		log.Printf("startsec")
	} else if typ == 9 {
		log.Printf("elementsec")
	} else if typ == 10 {
		veclen := U32(bufr)
		log.Printf("codesec veclen: %d", veclen)
		for veclen > 0 {
			size := U32(bufr)
			log.Printf("codesec size: %d", size)
			nlocals := U32(bufr)
			log.Printf("codesec nlocals: %d", nlocals)
			for nlocals > 0 {
				n := U32(bufr)
				log.Printf("codesec n: %d", n)
				t := U32(bufr)
				log.Printf("codesec t: %d", t)
				nlocals -= 1
			}
			doExpr(bufr)
			veclen -= 1
		}
	} else if typ == 11 {
		log.Printf("datasec")
		veclen := U32(bufr)
		log.Printf("datasec len: %d", veclen)
		for veclen > 0 {
			memidx := U32(bufr)
			log.Printf("datasec memidx: %d", memidx)
			doExpr(bufr)
			vecb := U32(bufr)
			log.Printf("datasec vecb: %d", vecb)
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
