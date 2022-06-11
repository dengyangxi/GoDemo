/****	接口继承实现	****/

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
	//初始化接口
	var iMy iMyInterface
	//初始化 结构
	mystruct := new(myStruct)
	iMy = mystruct
	mystruct.name = "张三"
	mystruct.age = 18

	iMy.Hello(mystruct)

	mystruct.name = "李四"
	mystruct.age = 20
	iMy.Hi(" 打招呼呀")
}

//定义一个接口
type iMyInterface interface {
	//打个招呼    hi ....
	Hi(interface{})

	//打个招呼    hello ....
	Hello(interface{})
}

//定义一个结构
type myStruct struct {
	name string
	age  int
}

//实现Hi 方法的业务逻辑
func (m *myStruct) Hi(i interface{}) {
	fmt.Println(" Hi  姓名：", m.name, "年龄：", m.age, "参数：", i)
}

//实现Hello 方法的业务逻辑
func (m *myStruct) Hello(interface{}) {
	fmt.Println(" Hello 姓名：", m.name, "年龄：", m.age)
}
