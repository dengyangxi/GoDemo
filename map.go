/****	数组	****/

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

	//map 在使用时需要使用 make 进行主动创建
	var myMap = make(map[string]string)
	//往map集合添加值
	myMap["one"] = "第一个值"
	myMap["two"] = "第二个值"
	myMap["three"] = "第三个值"

	//根据key删除一个值
	delete(myMap, "two")

	//从map集合中获取一个key的value值
	v, b := myMap["three"]
	if b {
		fmt.Println("您获取到的map集合的值:", v)
	} else {
		fmt.Println("不好意思,未能获取到map集合的值")
	}
	//循环输出 map集合的值
	for i, v := range myMap {
		fmt.Printf("Map集合 Index: [%v], Value: [%v]\n", i, v)
	}
}
