package parser

import (
	"bufio"
	"log"
)

func Name(bufr *bufio.Reader) string {
	nmlen := U32(bufr)
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

func Limit(bufr *bufio.Reader) {
	b, err := bufr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	n := U32(bufr)
	if int(b) == 0 {
		log.Printf("limit %d", n)
	} else {
		m := U32(bufr)
		log.Printf("limit %d %d", n, m)
	}
}
