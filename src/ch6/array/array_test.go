package array

import "testing"

func TestArray(t *testing.T) {
	var arr [5]int
	t.Log(arr)
	arr2 := [5]int{1, 2, 3, 4, 5}
	t.Log(arr2)
	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(arr3)
}

func TestArray2(t *testing.T) {
	a := [...]string{"Jack", "Tom", "Tony", "Kim"}
	b := a
	b[0] = "Tim"
	t.Log("a : ", a)
	t.Log("b : ", b)
}
