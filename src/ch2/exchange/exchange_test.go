package exchange

import "testing"

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	a, b = b, a

	t.Log(a)
	t.Log(b)
}
