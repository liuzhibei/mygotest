package main

import "fmt"

//数组定义
func arrayone() {
	//定义数组
	var arrone [3]int
	//定义类型并赋值
	var arrtwo [2]int = [2]int{1, 2}
	//编译器决定类型1
	var arrthree = [4]int{4, 4, 3, 3}
	//编译器决定类型2
	var arrfour = [...]int{11, 22}
	//简短方式
	arrfive := [2]int{3, 4}
	//二维数组定义
	var grid [2][3]int
	//二位数组赋值
	var gridtwo = [2][4]int{{22, 22, 33, 44}, {11, 22, 44, 44}}
	fmt.Println(arrone, arrtwo, arrthree, arrfour, arrfive, grid, gridtwo)
}

//数组遍历
func arrayprint(arr [3]string) {
	//用len遍历只能通过下标取值
	for i := 0; i < len(arr); i++ {
		fmt.Println("len:", i, " => ", arr[i])
	}
	//用range遍历，可以同事拿到下标和值
	for i, v := range arr {
		fmt.Println("range:", i, " => ", v)
	}
}

//数组为值传参，无法修改外部数组
func updatearray(arr [3]int) {
	arr[2] = 100
}

//指针数组传参，外部数组可被修改
func updatearraypoint(arr *[3]int) {
	arr[2] = 200
}

//数组传参
func paramarray() {
	arr := [3]int{1, 2, 3}
	//值传递
	updatearray(arr)
	fmt.Println("数组值传递无法被修改: ", arr)
	//指针传递
	updatearraypoint(&arr)
	fmt.Println("指针数组传递可修改：", arr)

}

func main() {
	//数组定于与赋值
	arrayone()
	//数组遍历
	arrayprint([3]string{"aaa", "bbb", "ccc"})
	//数组传参
	paramarray()
}
