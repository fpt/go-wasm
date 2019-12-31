package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type WAType int

const (
	I32 WAType = iota
	I64 WAType = iota
	F32 WAType = iota
	F64 WAType = iota
)

func doValType(b byte) WAType {
	switch b {
	case 0x7e:
		return I64
	case 0x7d:
		return F32
	case 0x7c:
		return F64
	}
	// 0x7f
	return I32
}

// u32 LEB128 decoder
func doU32(bufr *bufio.Reader) uint32 {
	return uint32(doLebUInt(bufr, 32))
}
func doLebUInt(bufr *bufio.Reader, nn int) int {
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	n := int(b)
	if n < 0x7F && n < 2^nn {
		return n
	}
	m := doLebUInt(bufr, nn-7)
	fmt.Printf(" m = %d\n", m)
	n = 0x7F*m + n - 0x7F
	return n
}
func doValTypeVec(bufr *bufio.Reader) uint32 {
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
		doValType(b)
	}
	return uint32(1) + cval
}
func doFunc(bufr *bufio.Reader) uint32 {
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
	c += doValTypeVec(bufr)
	c += doValTypeVec(bufr)
	return c
}
func doName(bufr *bufio.Reader) string {
	nmlen := doU32(bufr)
	name := make([]byte, nmlen)
	n, err := bufr.Read(name)
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	if n < int(nmlen) {
		log.Fatalf("File too short!")
	}
	return string(name)
}
func doLimit(bufr *bufio.Reader) {
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	n := doU32(bufr)
	if int(b) == 0 {
		log.Printf("limit %d", n)
	} else {
		m := doU32(bufr)
		log.Printf("limit %d %d", n, m)
	}
}
func doImport(bufr *bufio.Reader) {
	mod := doName(bufr)
	nm := doName(bufr)
	log.Printf("import mod: %s, name: %s", mod, nm)
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	typ := int(b)
	switch typ {
	case 0x00:
		idx := doU32(bufr)
		log.Printf("typeidx %d", idx)
	case 0x01:
		b, err := bufr.ReadByte()
		if err != nil {
			log.Fatalf("Error occured %s", err)
		}
		doLimit(bufr)
		log.Printf("tabletype %d", int(b))
	case 0x02:
		doLimit(bufr)
		log.Printf("memtype")
	case 0x03:
		b, err := bufr.ReadByte()
		if err != nil {
			log.Fatalf("Error occured %s", err)
		}
		typ := doValType(b)
		b, err = bufr.ReadByte()
		if err != nil {
			log.Fatalf("Error occured %s", err)
		}
		mut := int(b)
		log.Printf("globaltype typ:%v mut:%d", typ, mut)
	}
}
func doSkip(bufr *bufio.Reader, len int) {
	cont := make([]byte, len)
	n, err := bufr.Read(cont)
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	if n < len {
		log.Fatalf("File too short!")
	}
}
func doSection(bufr *bufio.Reader) {
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	typ := int(b)
	if typ == 1 {
		seclen := doU32(bufr)
		log.Printf("typesec len: %d", seclen)
		veclen := doU32(bufr)
		log.Printf("funcvec len: %d", veclen)
		for seclen > 0 && veclen > 0 {
			seclen -= doFunc(bufr)
			veclen -= 1
		}
	} else if typ == 2 {
		seclen := doU32(bufr)
		log.Printf("importsec len: %d", seclen)
		veclen := doU32(bufr)
		log.Printf("importvec len: %d", veclen)
		for veclen > 0 {
			doImport(bufr)
			veclen -= 1
		}
	} else if typ == 3 {
		log.Printf("funcsec")
	} else if typ == 4 {
		log.Printf("tablesec")
	}
}
func main() {
	file, err := os.Open("./wasm/a.out.wasm")
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	defer file.Close()
	/*stats, statsErr := file.Stat()
	 if statsErr != nil {
		log.Fatalf("Error occured %s", statsErr)
	 }
	 var size int64 = stats.Size()
	bytes := make([]byte, size)
	*/
	bufr := bufio.NewReader(file)
	var n int
	hdr := make([]byte, 4)
	ver := make([]byte, 4)
	n, err = bufr.Read(hdr)
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	if n < 4 {
		log.Fatalf("File too short!")
	}
	if !(hdr[0] == 0 && string(hdr[1:]) == "asm") {
		log.Fatalf("Not a wasm file.")
	}
	fmt.Println(string(hdr[1:]))
	n, err = bufr.Read(ver)
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	if n < 4 {
		log.Fatalf("File too short!")
	}
	doSection(bufr)
	doSection(bufr)
	return
}
