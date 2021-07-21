package main

import (
	"fmt"
	"math"
)

//包内部变量（go语言没有全局变量）
var aa, bb int
var cc string = "string"
var dd, ee = 10, "kfjdkf"

//使用（）集中定义
var (
	ff = "jjj"
	gg = 1000
)

//ff := 200 函数外无法使用：=

//go变量只声明不赋值的话会默认为0值
func variableZeroValue() {
	var i int
	var s string
	var b bool
	fmt.Printf("%d, %q, %t\n", i, s, b)
}

//变量初始化赋值
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "hahaha"
	fmt.Println(a, b, s)
}

//变量初始化赋值可以省略变量类型，变量根据值找到类型
func variableTypeDeduction() {
	var a, b, s, t = 10, 20, "kjfkdj", true
	fmt.Println(a, b, s, t)
}

//函数内可以使用：=定义并赋值变量（注意：变量不能重复定义）
func variableShorter() {
	a, s, t := 20, "kdjfk", false
	//a := 100 重复定义会报错
	fmt.Println(a, s, t)
}

//go语言只有强制类型转换，没有隐式类型转换
func forceTranceVariable() {
	var a, b int = 3, 4
	var c int
	var d float64
	// c = int(math.Sqrt(a*a + b*b)) Sqrt()函数参数必须是float64类型
	c = int(math.Sqrt(float64(a*a + b*b)))
	d = math.Sqrt(float64(a*a + b*b))
	fmt.Println(c, d)
}

func main() {
	//变量初始化零值
	variableZeroValue()
	//变量赋值
	variableInitialValue()
	//变量类型自动获取
	variableTypeDeduction()
	//变量简化定义赋值
	variableShorter()
	//包变量
	fmt.Println(aa, bb, cc, dd, ee, ff, gg)
	//强制类型转换
	forceTranceVariable()
}
