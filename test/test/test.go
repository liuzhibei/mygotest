package main

//可以使用vscode工具自动生成测试文件，然后实现表格驱动测试
func add(a, b int) int {
	return a + b
}

func ab(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
