package parser

import (
	"bufio"
	"log"
)

func doExpr(bufr *bufio.Reader) {
	for {
		instr, err := bufr.ReadByte()
		if err != nil {
			log.Fatalf("Error occured %s", err)
		}
		switch instr {
		case 0x00:
			log.Printf("unreachable")
		case 0x01:
			log.Printf("nop")
		case 0x02:
			log.Printf("block")
		case 0x03:
			log.Printf("loop")
		case 0x04:
			log.Printf("if")
		case 0x0B:
			log.Printf("end")
			return
		case 0x0C:
			log.Printf("br")
		case 0x0D:
			log.Printf("br_if")
		case 0x0E:
			log.Printf("br_table")
		case 0x0F:
			log.Printf("return")
		case 0x10:
			log.Printf("call")
			funcidx := U32(bufr)
			log.Printf("call funcidx: %d", funcidx)
		case 0x11:
			log.Printf("call_indirect")
		case 0x20:
			log.Printf("local.get")
		case 0x21:
			log.Printf("local.set")
		case 0x22:
			log.Printf("local.tee")
		case 0x23:
			log.Printf("global.get")
		case 0x24:
			log.Printf("global.set")
		case 0x28:
			log.Printf("i32.load")
		case 0x29:
			log.Printf("i64.load")
		case 0x2A:
			log.Printf("f32.load")
		case 0x2B:
			log.Printf("f64.load")
		case 0x2C:
			log.Printf("i32.load8_s")
		case 0x2D:
			log.Printf("i32.load8_u")
		case 0x2E:
			log.Printf("i32.load16_s")
		case 0x2F:
			log.Printf("i32.load16_u")
		case 0x30:
			log.Printf("i64.load8_s")
		case 0x31:
			log.Printf("i64.load8_u")
		case 0x32:
			log.Printf("i64.load16_s")
		case 0x33:
			log.Printf("i64.load16_u")
		case 0x34:
			log.Printf("i64.load32_s")
		case 0x35:
			log.Printf("i64.load32_u")
		case 0x36:
			log.Printf("i32.store")
		case 0x37:
			log.Printf("i64.store")
		case 0x38:
			log.Printf("f32.store")
		case 0x39:
			log.Printf("f64.store")
		case 0x3a:
			log.Printf("i32.store8")
		case 0x3b:
			log.Printf("i32.store16")
		case 0x3c:
			log.Printf("i64.store8")
		case 0x3d:
			log.Printf("i64.store16")
		case 0x3e:
			log.Printf("i64.store32")
		case 0x3f:
			log.Printf("memory.size")
		case 0x40:
			log.Printf("memory.grow")
		case 0x41:
			log.Printf("i32.const")
			n := U32(bufr)
			log.Printf("i32.const n: %d", n)
		case 0x42:
			log.Printf("i64.const")
			n := U32(bufr)
			log.Printf("i64.const n: %d", n)
		case 0x43:
			log.Printf("f32.const")
		case 0x44:
			log.Printf("f64.const")
		// i32 cmp
		case 0x45:
			log.Printf("i32.eqz")
		case 0x46:
			log.Printf("i32.eq")
		case 0x47:
			log.Printf("i32.ne")
		case 0x48:
			log.Printf("i32.lt_s")
		case 0x49:
			log.Printf("i32.lt_u")
		case 0x4A:
			log.Printf("i32.gt_s")
		case 0x4B:
			log.Printf("i32.gt_u")
		case 0x4C:
			log.Printf("i32.le_s")
		case 0x4D:
			log.Printf("i32.le_u")
		case 0x4E:
			log.Printf("i32.ge_s")
		case 0x4F:
			log.Printf("i32.ge_u")
		// i64 cmp
		case 0x50:
			log.Printf("i64.eqz")
		case 0x51:
			log.Printf("i64.eq")
		case 0x52:
			log.Printf("i64.ne")
		case 0x53:
			log.Printf("i64.lt_s")
		case 0x54:
			log.Printf("i64.lt_u")
		case 0x55:
			log.Printf("i64.gt_s")
		case 0x56:
			log.Printf("i64.gt_u")
		case 0x57:
			log.Printf("i64.le_s")
		case 0x58:
			log.Printf("i64.le_u")
		case 0x59:
			log.Printf("i64.ge_s")
		case 0x5A:
			log.Printf("i64.ge_u")
		// f32 cmp
		case 0x5B:
			log.Printf("f32.eq")
		case 0x5C:
			log.Printf("f32.ne")
		case 0x5D:
			log.Printf("f32.lt")
		case 0x5E:
			log.Printf("f32.gt")
		case 0x5F:
			log.Printf("f32.le")
		case 0x60:
			log.Printf("f32.ge")
		// f64 cmp
		case 0x61:
			log.Printf("f32.eq")
		case 0x62:
			log.Printf("f32.ne")
		case 0x63:
			log.Printf("f32.lt")
		case 0x64:
			log.Printf("f32.gt")
		case 0x65:
			log.Printf("f32.le")
		case 0x66:
			log.Printf("f32.ge")
		}
	}
}
