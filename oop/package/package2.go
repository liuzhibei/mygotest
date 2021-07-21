package pack

//结构体函数可定义在同一包下不同文件
func (mystruct *MyStruct) GetAge() int {
	return mystruct.Age
}
