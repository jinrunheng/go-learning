package operator

import (
	"fmt"
	"testing"
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 3 // 0011
	t.Log("Readable", a&Readable == Readable)
	t.Log("Writable", a&Writable == Writable)
	t.Log("Executable", a&Executable == Executable)

	// set unreadable
	fmt.Println("set unreadable")
	a = a &^ Readable

	t.Log("Readable", a&Readable == Readable)
	t.Log("Writable", a&Writable == Writable)
	t.Log("Executable", a&Executable == Executable)

	// set executable
	fmt.Println("set executable")
	a = a | Executable

	t.Log("Readable", a&Readable == Readable)
	t.Log("Writable", a&Writable == Writable)
	t.Log("Executable", a&Executable == Executable)
}
