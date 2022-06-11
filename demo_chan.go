/****	chan 通道	****/

//运行方式 cmd命令： go run "e:\项目\MyGO\inherit.go"
//	      VS Code： 右击 run code
//		  GoLand :  点击菜单栏的运行按钮
//程序所在包
//注意：启动程序必须在 main 包里面
package main

import "fmt"

//包引用
//类似C#  using fmt

//主要启动方法
//注意： 在go语言当中  方法、结构、接口、等等.. 大括号 { 必须和 方法名称 是一行，否则会报错
//       比如：  func main() {     // 编译通过
//  			 func main()      // 编译失败
//				 {			      // 大括号 和 方法名称 不再一行{
func main() {

	//定义一个 整型 通道
	var myChan = make(chan int)
	// 开启一个并发匿名函数  go 协程
	// 如果此处 不开启协程 所有的 goroutine（包括main）都处于等待 goroutine。也就是说所有 goroutine 中的 channel 并没有形成发送和接收对应的代码。
	// 即： 如果你将 go func(i int) {    改成(移除go)   func(i int) {    将会报错：  fatal error: all goroutines are asleep - deadlock!
	go func(i int) {
		// 往通道里面添加值
		myChan <- 100
		myChan <- 200
		myChan <- 300
		myChan <- i
		//关闭通过通道
		close(myChan)
	}(400)

	//关闭通过通道
	close(myChan)
	for v := range myChan {
		fmt.Println("获取到的值：", v)
	}

}
