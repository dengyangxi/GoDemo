package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Person
	school string
}

func (Person) SayHello(name string) string {
	fmt.Println(name + "这是来自 person的 SayHello。this is from Person")
	return "334"
}

func main() {

	stu := Student{school: "middle"}
	stu.name = "leo"
	stu.age = 30
	fmt.Print(stu.name)

	var aa = stu.SayHello("张三同学")
	fmt.Println(aa)

	//======================================================================
	//1. 数组的基本操作
	/* 定义一个长度为 3的数组  即： [3] */
	var sz = [3]string{"one", "two", "three"}
	/* 定义一个让编译器自动推算长度的数组 即：[...] */
	/*  var sz2 = xxx  简写  sz2 :=  xxx  */
	sz2 := [...]int{1, 2, 3, 4, 5, 6, 7}

	/*数组取值，并打印*/
	fmt.Println(sz[1])
	fmt.Println(sz2[6])

	/*循环打印数组*/
	for i := 0; i < len(sz); i++ {
		fmt.Println(sz[i])
	}

	/*循环打印数组，使用占位符*/
	for i := 0; i < len(sz2); i++ {
		fmt.Printf("Index:[%v],Value:[%b]\n", i, sz2[i])
	}

	//======================================================================
	var mymap = make(map[string]string)
	mymap["one"] = "第一个值"
	mymap["two"] = "第二个值"
	mymap["three"] = "第三个只"

	//删除 一个值
	delete(mymap, "two")

	/* 从map中获取一个key的Value */
	v, o := mymap["two"]
	/* o 表示是否获取到值，  v 表示获取的值内容 */
	if o {
		fmt.Println(v)
	} else {
		fmt.Println(o)
	}
	/* 使用 range 函数 变量 数组map */
	for k, v := range mymap {
		fmt.Printf("Key : [%v] , Value : [%v] \n", k, v)
	}

	//======================================================================

	//带有线程锁，  类似于C#中的 ConcurrentDictionary 线程自带线程安全
	var syncMap sync.Map
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")
	syncMap.Store("key3", "value4")

	v1, o2 := syncMap.Load("key2")
	if o2 {
		// map中存在key2
		fmt.Println(v1)
	} else {
		//map中不存在key2,  v1的值是： nil
		fmt.Println(o2)
	}

	//======================================================================

	/* 匿名函数-定义后立即调用 */
	func(data int) {
		fmt.Println("匿名函数打印值：", data)
	}(999)

	/* 匿名函数 赋值给变量 */
	fn := func(str string) {
		fmt.Println("您调用了变量匿名函数fn :", str)
	}
	//调用这个变量 匿名函数
	fn("张三")

	//将变量匿名函数 当做参数传递给 其他方法执行
	myFn("李四", fn)

	//======================================================================
	var invoker Invoker
	s := new(mystruct)
	invoker = s

	invoker.Call("hello word")
	invoker.Bak("Say Hi")

	s.age = 18
	s.name = "张三"

	s.Call("hello word")
	s.Bak("Say Hi")

	//======================================================================
	/* 结构的 实例化 和 赋值  一共有如下三种 */
	var mystruct1 mystruct
	mystruct1.age = 19
	mystruct1.name = "李四"
	fmt.Println(mystruct1.name)

	mystruct2 := &mystruct{age: 20, name: "王五"}
	fmt.Println(mystruct2.name)

	mystruct3 := new(mystruct)
	mystruct3.age = 21
	mystruct3.name = "小明"
	fmt.Println(mystruct3.name)

	//======================================================================
	/* 先执行一个方法，再判断结果 */
	if result := CheckData(89); result {
		fmt.Println("结果为true")

	} else {
		fmt.Println("结果为false")
	}
	// 上门方法 等价于：
	if CheckData(89) == true {
		fmt.Println("结果为:true")

	} else {
		fmt.Println("结果为: false")
	}

	//======================================================================
	//定义一个切片
	c := make(chan int)
	go func() {
		c <- -1
		c <- 2
		c <- 3
		close(c)
	}()
	/* 循环遍历 */
	for v := range c {
		fmt.Println(v)
	}

	//======================================================================
	/*  闭包 */
	var f func(int) bool
	f = CheckData
	f(18)

	//======================================================================
	//匿名行数赋值给变量
	closefunc := closeFn("张三", 18)
	//执行函数
	name, year := closefunc()
	fmt.Println("姓名："+name, "出生年份:", year)

	//======================================================================
	/* 延迟语句   defer  有点儿像 C# 里面的 finally */
	valueLock.Lock()
	defer valueLock.Unlock()
	valueKey := valueByKey["test"]
	fmt.Println(valueKey)

	//======================================================================
	/*  go 协程 */
	//计数器
	wg.Add(1)
	go mygo(151)
	fmt.Println("==========")
	//等等协程 执行
	wg.Wait()

	//======================================================================
	/*  协程 通讯 channel */
	//管道先进先出
	var chan1 = make(chan int, 3)
	//放数据到管道
	chan1 <- 10
	chan1 <- 20
	chan1 <- 30
	//从管道里面读取数据
	x := <-chan1
	fmt.Println(x)
	//必须要关闭 管道，才能 range 循环
	close(chan1)

	for ccc := range chan1 {
		fmt.Println("遍历管道：", ccc)
	}

	//======================================================================
	/*  协程 通讯 channel */
	channel := make(chan int, 100)
	channel2 := make(chan int, 100)

	go fnChannel(channel)
	go fnChannel2(channel, channel2)

	for channel := range channel2 {
		fmt.Println("channel2 取值：", channel)
	}
}

func fnChannel(channel chan int) {
	for i := 0; i < 1000; i++ {
		channel <- i
	}
	close(channel)
}

func fnChannel2(channel chan int, channel2 chan int) {
	//遍历通道channel
	for v := range channel {
		fmt.Println("中转站 ：", v)
		//将通道 1 取到的值 +100 存入通道 channel2
		channel2 <- v + 100
	}
	close(channel2)
}

var wg sync.WaitGroup

func mygo(i int) {
	fmt.Println("协程执行方法", i)
	//计数器完成一次   即：计数器 -1 ，对应 wg.Add(1)
	wg.Done()
}

var (
	valueByKey = make(map[string]interface{})

	// 保证安全的互斥锁
	valueLock sync.Mutex
)

// 定义一个返回函数的方法，然后内部使用闭包实现
func closeFn(name string, age int) func() (string, int) {
	year := time.Now().Year() - age
	//匿名方法
	return func() (string, int) {
		return name, year
	}
}

func CheckData(a int) bool {
	fmt.Println("CheckData 请求参数：", a)
	if (a % 2) == 1 {
		return true
	} else {
		return false
	}
}

///执行匿名函数
func myFn(str string, fn func(string)) {
	fn(str)
}

type Invoker interface {
	Call(interface{})
	Bak(interface{})
}

type mystruct struct {
	name string
	age  int
}

func (s *mystruct) Call(p interface{}) {

	fmt.Println("MyStruct ", p, s.age, s.name)
}
func (s *mystruct) Bak(p interface{}) {
	fmt.Println("MyStruct ", p, s.age, s.name)
}
