package _type

import "testing"

func TestString(t *testing.T)  {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
