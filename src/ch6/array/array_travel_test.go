package array

import (
	"testing"
)

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}

	// 1
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	// 2
	for idx, e := range arr {
		t.Log(idx, e)
	}

	// 3
	for _, e := range arr {
		t.Log(e)
	}
}
