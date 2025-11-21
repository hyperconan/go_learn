package test_closure

import "fmt"

type Person struct {
	Name string `json:"name"`
}

func (p *Person) marry(lover string) string {
	return p.Name + " is married to " + lover
}

var i = 4

func Run() {
	fmt.Print("test_closure")

	returnFunc := func() func(int, string) (int, string) {
		fmt.Println("this is a anonymous function,没有入参的")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}() // 声明匿名函数且立即执行

	// 执行returnFunc闭包并传递参数
	ret1, ret2 := returnFunc(1, "test")
	fmt.Printf("ret1: %d, ret2: %s\n", ret1, ret2)

	//声明一个变量，变量类型是返回string的，入参是string的函数类型
	var callsomeone func(string) string
	p := Person{"conan"}
	callsomeone = p.marry
	fmt.Println(callsomeone("love"))

	if i := 3; i < 5 {
		fmt.Println(i)
	}
	fmt.Println(i)
}
