package main

import "fmt"

//创建map
func mymap() {
	var m map[string]int
	var m1 map[string]string = map[string]string{
		"aaa": "aaaaaaa",
		"bbb": "bbbbbbb",
	}
	var m2 = map[string]string{
		"ccccc": "ccccccc",
		"ddddd": "ddddddd",
	}
	m3 := map[string]string{
		"hahaha": "hahah",
		"hehehe": "hehehe",
	}
	m4 := make(map[string]string)
	fmt.Println(m, m1, m2, m3, m4)
}

//获取key的值
func getkeyvalue() {
	m := map[string]string{
		"a": "11",
		"b": "22",
		"c": "33",
	}
	//a := m["aa"] 拿不到会返回空
	if a, ok := m["aa"]; ok {
		fmt.Println(a)
	} else {
		fmt.Println("key is not exists")
	}
}

func delmapkey() {
	m := map[string]string{
		"a": "11",
		"b": "22",
		"c": "33",
	}
	delete(m, "a")
	fmt.Println(m)
}

//遍历map
func traverse(mm map[string]string) {
	for key, val := range mm {
		fmt.Println(key, val)
	}
}

func main() {
	//创建map
	mymap()
	//获取map元素值
	getkeyvalue()
	//删除元素
	delmapkey()
}
