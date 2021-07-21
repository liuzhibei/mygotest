package main

import (
	"fmt"
	"reflect"
	"runtime"
)

//两个参数，一个返回值
func add(a, b int) int {
	return a + b
}

//两个参数，两个返回值
func div(a, b int) (int, int) {
	return a / b, a % b
}

//返回值命名
func returnname(a, b int) (c, d int) {
	c = b
	d = a
	return
}

//go函数第二个参数返回错误
func returnerr(a, b int) (bool, error) {
	if a > b {
		return true, nil
	} else {
		return false, fmt.Errorf("error: %d < %d", a, b)
	}
}

//函数做参数(函数为一等公民)
func paramfunc(op func(int, int) int, a, b int) int {
	//通过映射获取函数名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling func %s with args %d, %d\n", opName, a, b)
	return op(a, b)
}

//可变参数列表（可传多参数）
func sumnum(numbers ...int) int {
	for key, val := range numbers {
		//获取键与值
		fmt.Println(key, val)
	}
	var sum int
	for i := range numbers {
		//通过健访问变量
		sum += numbers[i]
	}
	return sum
}

func main() {
	//函数一个返回值
	fmt.Println(add(10, 5))
	//函数两个返回值
	fmt.Println(div(13, 3))
	//函数返回值命名（一般不用，可读性不好）
	fmt.Println(returnname(10, 20))
	//函数返回错误信息
	res, err := returnerr(10, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	//函数参数(传函数名称就可以了)
	fmt.Println(paramfunc(add, 1, 3))
	//函数参数（匿名函数做参数）
	fmt.Println(paramfunc(
		func(a, b int) int {
			return a * b
		}, 10, 3))
	//可变参数列表
	fmt.Println(sumnum(10, 20, 30, 40))
}
