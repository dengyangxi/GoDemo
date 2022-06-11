package main

import (
	"fmt"
	"sync"
)

func main() {
	/* 延迟语句   defer  有点儿像 C# 里面的 finally */
	valueByKey["name"] = "simon"
	//锁定资源
	valueLock.Lock()
	//解锁
	defer valueLock.Unlock()
	valueByKey["two"] = "第二个value"
	valueByKey["three"] = "第三个value"
	valueKey := valueByKey["two"]
	fmt.Println(valueKey)

	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")

	//延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。
	// 使用 defer  修饰的代码，  就相当于C#  finally 内的代码， 不管执行成功 还是失败 都会执行

}

var (
	valueByKey = make(map[string]interface{})

	//保证安全的互斥锁
	valueLock sync.Mutex
)
