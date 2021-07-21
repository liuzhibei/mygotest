package main

import (
	"fmt"
)

//定义struct
type node struct {
	val         string
	left, right *node
}

//struct赋值
func mystruct() {
	//创建struct
	structone := node{}
	//new返回的是struct的地址
	structtwo := new(node)
	fmt.Println(structone, structtwo)
	//赋值部分属性时要有key
	structthree := node{val: "hha"}
	//赋值所有时可省略key
	structfour := node{"hehe", nil, nil}
	fmt.Println(structthree, structfour)
}

//访问struct元素要使用.
func visitstruct() {
	structone := node{}
	structone.val = "111"
	structone.left = createnode("222")
	structone.right = createnode("333")
	fmt.Println(structone.left.right)
}

//工厂函数
func createnode(val string) *node {
	//返回了局部变量的地址（在go语言中可以使用）
	return &node{val: val}
}

//结构体定义方法
func (nodeone *node) traverse() {
	if nodeone == nil {
		return
	}
	fmt.Println(nodeone.val)
	nodeone.left.traverse()
	nodeone.right.traverse()
}
func main() {
	//创建struct
	mystruct()
	//访问struct的属性
	visitstruct()
	//struct的工厂函数
	fmt.Println(createnode("create"))
	//调用结构体方法，如果需要修改外部数据，需要指针。
	structfunc := node{
		"111",
		&node{"222", nil, nil},
		&node{val: "333"},
	}
	//调用指针方法会自动转化，不用显示指针访问
	structfunc.traverse()
}
