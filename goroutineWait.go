package main

import (
	"fmt"
	"sync"
)

func main() {

	//设置计数器 总数量
	wg.Add(10)

	//执行协程，传入10个协程
	go mygo(10)

	//等等协程全部完成
	wg.Wait()

	//请确保 计数器总数量 = Done() 执行次数，

	//如果计数器总数量是 10次， 即： wg.Add(10)  但是 wg.Done() 执行次数 小于 10次  将会出现如下死锁错误：
	//fatal error: all goroutines are asleep - deadlock!

	//如果计数器总数量是 10次， 即： wg.Add(10)  但是 wg.Done() 执行次数 大于 10次  将会出现如下错误：
	//panic: sync: negative WaitGroup counter

	fmt.Println("Done")
}

func mygo(data int) {
	for i := 0; i < data; i++ {
		//打印数据
		fmt.Println("循环打印：", i)
		//计数器 完成一次
		wg.Done()
	}

}

var (

	//计数器
	wg = sync.WaitGroup{}
)
