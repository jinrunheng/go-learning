package condition

import "testing"

func TestIfMultiSec(t *testing.T) {
	a := 1
	b := 2
	if v, err := add(a, b); err == nil {
		t.Log(v)
	} else {
		t.Log(err)
	}
}

func add(a, b int) (v int, err error) {
	return a + b, nil
}
