package test_array_slice

import "fmt"

func test_arr() {
	var m1 [2]map[int]string = [2]map[int]string{ //  一个包含两个map元素的数组
		{1: "A"}, {2: "B"},
	}
	fmt.Println(m1)                         // [map[1:A] map[2:B]]
	fmt.Println(m1[1])                      // map[2:B]
	fmt.Println(m1[0])                      // map[1:A]
	var m2 map[int]string = map[int]string{ //  一个map元素
		1: "A", 2: "B",
	}
	fmt.Println(m2)    // map[1:A 2:B]
	fmt.Println(m2[1]) // A
	fmt.Println(m2[2]) // B
	fmt.Println(m2[0]) //测试空值
}

func Run() {
	test_arr()
}
