package _type

import "testing"

func TestPoint(t *testing.T)  {
	a := 1
	aPtr := &a
	t.Log(a,aPtr)
	t.Logf("%T %T",a,aPtr)
}