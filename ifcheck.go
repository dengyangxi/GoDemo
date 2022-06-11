/****	if 判断语句	****/

//运行方式 cmd命令： go run "e:\项目\MyGO\inherit.go"
//	      VS Code： 右击 run code
//		  GoLand :  点击菜单栏的运行按钮
//程序所在包
//注意：启动程序必须在 main 包里面
package main

//包引用
//类似C#  using fmt
import (
	"fmt"
)

//主要启动方法
//注意： 在go语言当中  方法、结构、接口、等等.. 大括号 { 必须和 方法名称 是一行，否则会报错
//       比如：  func main() {     // 编译通过
//  			 func main()      // 编译失败
//				 {			      // 大括号 和 方法名称 不再一行{
func main() {

	//先执行一个方法，在判断结果
	// 先执行： result := check(1) ;  在判断  result == 100
	if result := check(1); result == 100 {
		fmt.Println("符合判断条件，进入 if ")
	} else {
		fmt.Println("不符合判断条件， 进入 else")
	}

	//如下写法 和上面写法含义一样
	result := check(1)
	if result == 100 {
		fmt.Println("进入if判断")
	} else {
		fmt.Println("进入else")
	}
}

//定义一个方法
func check(i int) int {
	return i * 100
}
