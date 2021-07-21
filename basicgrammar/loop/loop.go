package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//for语句应用转换10进制数字为二进制数字
func converttobin(n int) string {
	var result string
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//相当于while
func printfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//for实现无线循环
func foreverprint() {
	for {
		fmt.Println("forever print")
	}
}

func main() {
	//for
	fmt.Println(converttobin(8))
	//for 实现while
	printfile("abc.txt")
	//for 实现无线循环
	foreverprint()
}
