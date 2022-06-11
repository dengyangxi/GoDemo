/****	继承	****/

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

	//实例化 Student ， 类似 C#  var stu = new Student(){ school: "上海市x小学",address: "徐汇区-漕宝路x号",   name="张三" , age =6}
	var stu = Student{school: "上海市x小学", address: "徐汇区-漕宝路x号"}
	//给学生属性赋值
	stu.name = "张三"
	stu.age = 6
	// 也可以使用 :=  和上面的语句意思的一样的   定义一个 stu变量  然后将赋值 Student
	//stu := Student{ school: "上海市x小学",address: "徐汇区-漕宝路x号"}

	//打印输出 学生姓名
	fmt.Println(stu.name)

	var result = stu.SayHi("李四")
	//打印输出调用 SayHi方法的返回结果
	fmt.Print(result)

}

//定义一个结构
//注意： struct 和 C# 一样是 值类型，由于go没有class 因此一般实体类 都定义成 struct
// 人的属性 （实体）
type Person struct {
	//姓名
	name string
	//年龄
	age int
}

//定义一个学生的 实体
type Student struct {
	// Student 继承 Person ,  类似C# 中语法  public class Student : Person
	Person

	//学校名称
	school string

	//学校地址
	address string
}

//定义一个SayHi方法
func (Person) SayHi(name string) string {
	//打印方法的请求参数信息
	fmt.Println("您调用了 SayHi方法,并传入了参数 name :", name)
	//将学生姓名 + 同学 称呼  返回
	return name + " 同学"
}
