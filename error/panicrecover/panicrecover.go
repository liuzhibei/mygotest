package main

import (
	"errors"
	"fmt"
)

//用recover捕获panic
func myrecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occured:", err)
		} else {
			panic(err)
		}
	}()
	panic(errors.New("this is an error"))
}

func main() {
	myrecover()
}
