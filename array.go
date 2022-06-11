/****	数组	****/

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

	fmt.Println("====写法 一：")
	//定义一个长度为3的数组，即： [3]
	var sz = [3]string{"one", "two", "three"}
	//数组取值，并打印。    注意： 数组 索引从 0 开始，和 C# 一样
	fmt.Println(sz[1])

	//循环打印数组
	for i := 0; i < len(sz); i++ {
		fmt.Printf("Index: [%v], Value: [%v]\n", i, sz[i])
	}

	fmt.Println("====写法 二：")
	//定义一个由编译器推算长度的数组。 即： [...]  ,代码编译的时候 会自动计算 数组的长度，替换掉...
	sz2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//循环打印数组
	for i, v := range sz2 {
		fmt.Printf("Index: [%v], Value: [%v]\n", i, v)
	}

}
