package main

import (
	"bufio"
	"fmt"
	"os"
)

//defer
func mydefer() {
	defer fmt.Println(1) //先进后出
	defer fmt.Println(2)
	fmt.Println(3) // 不会被panic，return打断
	defer fmt.Println(4)
	panic("error occurred")
}

//defer 用于关闭资源,刷新缓存
func closeresource(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file) //利用缓存加速写操作
	defer writer.Flush()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, i)
	}
}

func main() {
	//mydefer()
	closeresource("writer.txt")
}
