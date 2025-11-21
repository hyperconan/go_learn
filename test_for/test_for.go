package test_for

import "fmt"

func test_break() {
	switch i := 1; i {
	case 1:
		fmt.Println("进入case 1")
		if i == 1 {
			break //  break 直接调出整个switch 连default都不执行
		}
		fmt.Println("i等于1")
	case 2:
		fmt.Println("i等于2")
	default:
		fmt.Println("default case")
	}

tag: // 中断标记
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 7; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			//break tag // i=1 j=5
			continue tag // i=1,j=5 i=2,j=5 i=3,j=5
		}
	}
	fmt.Println("跳出循环后,继续执行")
}

func test_for() {
	b := 5
	for b < 10 { // 等同于 while b < 10 )
		fmt.Println(b)
	}

	s := []int{9, 7, 4, 76}
	for i, e := range s { // 等同于 enumerate(s)
		fmt.Println(i, e)
	}

	m := map[string]([]int){
		"a": []int{1, 2},
		"b": []int{2, 3},
	}
	for k := range m { // 等同于m.keys()
		fmt.Println(k)
	}
	for _, v := range m { // 等同于m.values()
		fmt.Println(v)
	}
	for k, v := range m { // 等同于 m.items()
		fmt.Println(k, v)
	}
}

func Run() {
	fmt.Println("test_for")
	//test_for()
	test_break()
}
