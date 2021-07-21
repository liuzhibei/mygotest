package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

//定义接口与方法
type retrier interface {
	get(url string) string
}

//定义第二个接口与方法
type poster interface {
	post(url string, form map[string]string) string
}

//接口组合
type retrierposter interface {
	retrier
	poster
}

//定义结构体
type retrierstruct struct {
	useragent string
	timeout   time.Duration
}

//结构体实现接口
func (r retrierstruct) get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}
	resp.Body.Close()

	return string(result)
}

func (r retrierstruct) post(url string, form map[string]string) string {
	return url
}

//接口使用者决定接口的方法
func download(r retrier) string {
	return r.get("http://www.baidu.com")
}

func main() {
	var r retrierposter //接口变量包含实现者的类型和实现者的指针指向实现者
	fmt.Printf("%T %v\n", r, r)
	r = retrierstruct{}
	fmt.Printf("%T %v\n", r, r)
	download(r)
}
