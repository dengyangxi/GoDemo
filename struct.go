/****	struct (model)实例化,复制	****/

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

	//一共三种实例化 赋值方式

	//初始化  写法一 ：
	var my_struct myStruct
	//给属性赋值 -  姓名
	my_struct.name = "张三"
	//给属性赋值- 年龄
	my_struct.age = 18
	//打印属性
	fmt.Println("写法一 姓名：", my_struct.name, " 年龄 :", my_struct.age)

	//初始化  写法二 ：
	my_struct2 := &myStruct{name: "李四", age: 20}
	//打印属性
	fmt.Println("写法二 姓名：", my_struct2.name, " 年龄 :", my_struct2.age)

	//初始化  写法三 ：
	my_struct3 := new(myStruct)
	//给属性赋值 -  姓名
	my_struct3.name = "王五"
	my_struct3.age = 22
	//打印属性
	fmt.Println("写法三 姓名：", my_struct3.name, " 年龄 :", my_struct3.age)
}

// 定义一个struct 结构， 类似 C# 中的 model实体类，因 go语言中没有 class，因此一般使用 struct
type myStruct struct {
	//姓名
	name string
	//年龄
	age int
}
