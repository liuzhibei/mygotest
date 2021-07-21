package main

import "fmt"

//参数值传递（go只有值传递， 没有引用传递）
func swag(a, b int) {
	b, a = a, b
}

// 指针传递（也是值传递，只不过指针是一个地址）
func swagpoint(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	a, b := 10, 15
	//参数值传递
	swag(a, b)
	fmt.Println(a, b)
	//参数指针传递
	swagpoint(&a, &b)
	fmt.Println(a, b)
}
