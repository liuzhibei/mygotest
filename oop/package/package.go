package pack

//大写开通是public
type MyStruct struct {
	Name string
	Age  int
}

//大写开头为public方法
func (mystruct *MyStruct) GetName() string {
	return mystruct.Name
}

//小写开头为private方法
func (mystruct *MyStruct) setName(name string) {
	mystruct.Name = name
}
