package main

import (
	"fmt"
	"io/ioutil"
)

//正常的if else语句
func myif() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

//if后可跟赋值语句+条件（赋值的变量作用域只能是if块）
func optimizeif() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Printf("%s\n", contents)
	} else {
		fmt.Println(err)
	}
	// fmt.Println(contents) if后的变量只能再if块使用
}

//switch语句
func switcheval(op string, a, b int) int {
	var result int
	switch op {
	case "+":
		result = a + b //不用break，匹配成功会自动break
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("无法识别的操作" + op)
	}
	return result
}

//switch实现分数等级判定
func switchgrade(score int) string {
	var result string
	switch {
	case score < 60:
		result = "F"
	case score < 80:
		result = "B"
	case score < 100:
		result = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return result
}

func main() {
	//if
	myif()
	//if后跟赋值语句
	optimizeif()
	//switch
	fmt.Println(switcheval("+", 10, 20))
	//switch
	fmt.Println(switchgrade(101))
}
