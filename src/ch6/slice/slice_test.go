package slice

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	t.Log(a)
	t.Log(reflect.TypeOf(a))
}

func TestSlice2(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	var b []int
	b = a[1:4]
	t.Log(reflect.TypeOf(b))
	t.Log(b)
}

func TestSliceShareMemory(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	b[0] = 100
	t.Log(b)
	t.Log(a)
}

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := make([]int, 3, 5)
	t.Log(len(s1), cap(s1))
	t.Log(s1[0], s1[1], s1[2])
}

func TestSliceInit2(t *testing.T) {
	s0 := make([]int, 3, 5)

	for i := 0; i < 5; i++ {
		s0 = append(s0, i)
	}

	t.Log(len(s0), cap(s0))
	for _, e := range s0 {
		t.Log(e)
	}
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s),cap(s))
	}
}

