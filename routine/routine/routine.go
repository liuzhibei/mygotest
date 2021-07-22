package main

import (
	"fmt"
	"time"
)

//goroutine与main函数数据并发关系
func goroutine() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("hahah%d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond) //不加上此句，则当mygo函数执行完后会销毁所有数据，还没有等到打印就结束了
}

func goroutine2() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				// runtime.Gosched() //goroutine是非抢占式的，因此需要此函数去自主交出控制权（mygo（））不用自主交出控制权是因为由io操作会切换线程间调度
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func main() {
	goroutine()
}
