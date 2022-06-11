/****	线程安全Map	****/

//运行方式 cmd命令： go run "e:\项目\MyGO\inherit.go"
//	      VS Code： 右击 run code
//		  GoLand :  点击菜单栏的运行按钮
//程序所在包
//注意：启动程序必须在 main 包里面
package main

import (
	"fmt"
	"sync"
)

//包引用
//类似C#  using fmt

//主要启动方法
//注意： 在go语言当中  方法、结构、接口、等等.. 大括号 { 必须和 方法名称 是一行，否则会报错
//       比如：  func main() {     // 编译通过
//  			 func main()      // 编译失败
//				 {			      // 大括号 和 方法名称 不再一行{
func main() {
	//带有线程锁，  类似于C#中的 ConcurrentDictionary 线程自带线程安全
	var syncMap sync.Map
	//添加值
	syncMap.Store("one", "Value 1")
	syncMap.Store("two", "Value 2")
	syncMap.Store("three", "Value 3")
	syncMap.Store("three", 1)
	//根据key获取一个值
	v, b := syncMap.Load("two")
	//如果取到值 b= true  v= 值 ,否则 b= false  v =nil,
	if b {
		fmt.Println("获取到的值是：", v)
	} else {
		fmt.Println("未能获取到值")
	}

	//循环
	syncMap.Range(
		func(k, v interface{}) bool {
			fmt.Printf("%v:%v \n", k, v)
			return true
		})
	fmt.Println("结束")
}
