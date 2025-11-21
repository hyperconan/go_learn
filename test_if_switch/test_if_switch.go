package test_if_switch

import "fmt"

func test_if() {
	var a = 1
	if b, c := 998, "hello"; a == b {
		fmt.Println("a == b")
	} else if a != 1 {
		fmt.Println("a != b")
	} else {
		fmt.Println(b) // b只在if语句块中起作用
		fmt.Println(c) // c只在if语句块中起作用
	}
	//fmt.Println(b) // 超出定义域 提示undefined
}

func Run() {
	fmt.Println("test_if_switch")
	test_if()
}
