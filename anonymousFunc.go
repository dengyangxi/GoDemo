/****	匿名函数	****/

//运行方式 cmd命令： go run "e:\项目\MyGO\inherit.go"
//	      VS Code： 右击 run code
//		  GoLand :  点击菜单栏的运行按钮
//程序所在包
//注意：启动程序必须在 main 包里面
package main

import (
	"fmt"
	"time"
)

//包引用
//类似C#  using fmt

//主要启动方法
//注意： 在go语言当中  方法、结构、接口、等等.. 大括号 { 必须和 方法名称 是一行，否则会报错
//       比如：  func main() {     // 编译通过
//  			 func main()      // 编译失败
//				 {			      // 大括号 和 方法名称 不再一行{
func main() {
	//定义一个匿名函数，并赋值给变量
	var fn = func(name string) string {
		return name + "-同学"
	}
	//调用匿名函数
	var result = fn("f")
	//打印匿名函数的值
	fmt.Println(result)

	//定义一个匿名函数，并调用
	func(age int) {
		var year = time.Now().Year() - age
		fmt.Println("传入年龄：", age, "出生年份：", year)
	}(18)

	var f func(int) bool
	f = CheckData
	f(18)

}

func CheckData(i int) bool {
	if i%2 == 1 {
		fmt.Println(" 传入的是： 奇数")
		return true
	} else {
		fmt.Println(" 传入的是： 偶数")
		return false
	}

}
