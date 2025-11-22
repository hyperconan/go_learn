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

	var numarr = [3]int{1, 2, 3}
	fmt.Println(numarr)

	var multidimArr = [3][2]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	fmt.Println(multidimArr) // 三行两列数组

	noModify(multidimArr) // 传入的值是副本

	fmt.Println(multidimArr)
	withModify(&multidimArr) // 传入的是指针，修改的是原数组

	fmt.Println(multidimArr)
}

func noModify(arr [3][2]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = 0
		}
	}
}

func withModify(arr *[3][2]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = 0
		}
	}
}

func test_slice() {
	var s1 = []int{} // 切片没有长度
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1[1:2])
	var a1 = [3]int{1, 2, 3}
	fmt.Println(a1[1:2]) //从数组中创建一个切片

	var s2 = make([]int, 1, 3) // 创建一个长度为0，容量为3的切片
	fmt.Println(s2)
	fmt.Println("len:", len(s2), " cap:", cap(s2))
	s2 = append(s2, 3, 4, 5, 6, 7, 8, 9, 10) // 重新申请空间将旧数组内容赋值到新空间中，s2重新指向新数组
	fmt.Println(s2)
	fmt.Println("len:", len(s2), " cap:", cap(s2))

	/*
		 ... 展开切片,s2[2:]... 下标为2开始的元素
		append([]int{998, 999}, s2[2:]...) 含义为 生成一个切片，包含998,999, 向这个切片写入s2[2:]... 下标为2开始的元素
		s2则为 s2[:2] 下标为0,1的元素 加上 998,999 加上 s2[2:]... 下标为2开始的元素
	*/
	s2 = append(s2[:2], append([]int{998, 999}, s2[2:]...)...)
	fmt.Println(s2)
	fmt.Println("len:", len(s2), " cap:", cap(s2))

	var s3 = s2
	fmt.Println("before s2:", s2)
	fmt.Println("before s3:", s3)
	s3[0] = 1000
	fmt.Println("after s2:", s2)
	fmt.Println("after s3:", s3)

	var s4 = make([]int, 2, 4)
	fmt.Println("before copy s2,s4", s2, s4)
	copy(s4, s2)                      // (dst,src) s2 -> s4
	fmt.Println("after copy s4,", s4) // s4 复制了 s2 的前两个元素

	var s5 = make([]int, 4, 6)
	fmt.Println("before copy s4,s5", s4, s5)
	copy(s5, s4)
	fmt.Println("after copy s5,", s5)
}

func Run() {
	//test_arr()
	test_slice()
}
