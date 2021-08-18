package _type

import (
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	var c MyInt

	// 只能显示地进行类型转换
	b = int64(a)
	// 别名和原有类型也不能进行隐式类型转换
	c = MyInt(b)
	t.Log(b)
	t.Log(c)
}
