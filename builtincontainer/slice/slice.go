package main

import "fmt"

//数组获取切片
func sliceOne() {
	//定义一个长度为8的数组
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//获取切片
	s1 := arr[2:6]
	s2 := arr[:4]
	s3 := arr[3:]
	s4 := arr[:]
	fmt.Println(arr, s1, s2, s3, s4)
	// 切片参数传递，可改变外部切片
	updateSlice(s2)
	fmt.Println(s2, s4, arr)
}

//切片传参（可替代数组指针传参方案）
func updateSlice(slice []int) {
	slice[0] = 100
}

//slice直接创建
func sliceTwo() {
	var s []int
	fmt.Println("slice零值:", s) //slice的零值为nil
	for i := 0; i < 10; i++ {
		printSlice(s)
		s = append(s, i)
	}

	var s1 []int = []int{1, 2}
	var s2 = []int{1, 2, 3}
	s3 := []int{1, 2, 3, 4}
	s4 := make([]int, 16)
	s5 := make([]int, 10, 32)
	fmt.Println(s1, s2, s3)
	printSlice(s4)
	printSlice(s5)
}

//slice复制
func copyslice() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 10, 20)
	copy(s2, s1)
	fmt.Println("copys1->s2:", s1, s2)
}

//删除slice元素
func deleteslice() {
	slice := []int{2, 3, 4, 5}
	s2 := append(slice[:2], slice[3:]...) //go支持将slice...作为参数列表传入函数
	fmt.Println(s2)
}

func printSlice(s []int) {
	fmt.Printf("slice的len=%d, cap=%d\n", len(s), cap(s))
}

//切片继续被取切片（可向后扩展，不可向前扩展）
func reslice() {
	arr := [5]int{1, 2, 3, 4, 5}
	sliceone := arr[1:4]
	//slice可以向后扩展到数组的cap容量长度，看到sliceone没有但是arr有的元素，但不能向前扩展
	slicetwo := sliceone[2:4]
	fmt.Println(arr, sliceone, slicetwo)
}

//添加元素
func sliceappend() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	sliceone := append(slice, 10)    //原数组最后一个元素被修改
	slicetwo := append(sliceone, 11) //当添加后超出原数组cap范围，需要新开辟一个数组作为view的底层数组，原数组如果没有人引用会被垃圾回收
	slicethree := append(slicetwo, 12)
	fmt.Println(arr, slice, sliceone, slicetwo, slicethree)
}

func main() {
	//数组获取slice
	sliceOne()
	//创建slice
	sliceTwo()
	//slice重复获取
	reslice()
	//切片添加元素
	sliceappend()
	//复制
	copyslice()
	//删除
	deleteslice()
}
