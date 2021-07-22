package main

import (
	"bufio"
	"fmt"
	"os"
)

//错误处理
func closeresource(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err) //未知的错误用panic
		} else {
			fmt.Printf("%s, %s, %s \n", pathError.Op, pathError.Path, pathError.Err) //已知错误记录日志等操作
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, i)
	}
}

func main() {
	closeresource("writer.txt")
}
