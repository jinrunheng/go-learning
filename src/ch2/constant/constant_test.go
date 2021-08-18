package constant

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func TestConstant(t *testing.T) {
	t.Log(Monday)
	t.Log(Tuesday)
	t.Log(Wednesday)
	t.Log(Thursday)
	t.Log(Friday)
	t.Log(Saturday)
	t.Log(Sunday)
}
