package main

import "fmt"

//函数闭包
func adder() func(num int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	add := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}
}
