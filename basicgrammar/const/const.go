package main

import (
	"fmt"
	"math"
)

//枚举类型
const (
	one   = 1
	two   = 2
	three = 3
)

//自增枚举类型iota
const (
	java = iota
	_
	php
	ccc
	python
)

//自增实现特殊枚举 b, kb, mb, gb, tb, pb
const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

//函数常量（常量不指定类型，可认为是文本类型）
func myconst() {
	const a, b int = 3, 4
	// c := math.Sqrt(a*a + b*b) 常量指定类型
	c := math.Sqrt(float64(a*a + b*b))
	const aa, bb = 3, 4 //常量数值可作为各种类型使用
	cc := int(math.Sqrt(aa*aa + bb*bb))
	const (
		ee        = 10
		ff string = "kfdjfkdjfd"
	)
	fmt.Println(c, cc, ee, ff)
}

func main() {
	//常量
	myconst()
	//枚举
	fmt.Println(one, two, three)
	//自增枚举
	fmt.Println(java, php, ccc, python)
	//特殊枚举
	fmt.Println(b, kb, mb, gb, tb, pb)
}
