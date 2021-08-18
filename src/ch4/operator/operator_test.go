package operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4}

	t.Log("a == b", a == b)
	t.Log("a == c", a == c)

	// d := [...]int{1, 2, 3, 4, 5}
	//t.Log("a == d",a == d) // wrong
}
