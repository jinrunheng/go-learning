package constant

import "testing"

// 位的偏移
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant2(t *testing.T) {
	a := 3 // 0011
	t.Log("Readable", a&Readable == Readable)
	t.Log("Writable", a&Writable == Writable)
	t.Log("Executable", a&Executable == Executable)
}
