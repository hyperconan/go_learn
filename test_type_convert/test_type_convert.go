package test_type_convert

import (
	"fmt"
	"strconv"
)

func test_int2byte() {
	a := 4294967295
	b := byte(a)
	fmt.Println("a to b:", b) // 高位丢弃
}

func test_str_convert() {
	a := "你a"
	b := []byte(a)
	c := []rune(a)
	fmt.Println("a to byte:", b) // 你-3个字节 a-1个字节
	fmt.Println("a to rune:", c) // Unicode 码点

	fmt.Printf("bytes v: %v \n", string(b))
	fmt.Printf("runes v: %v \n", string(c))

	numStr := "110"                         // string类型
	int_instance, _ := strconv.Atoi(numStr) // 将字符串转数字 类似于 python中int()转换一个字符串 // int类型
	fmt.Printf("int_instance: %d \n", int_instance)
	i2s := strconv.Itoa(int_instance) // string 类型
	fmt.Println("i2s:", i2s)          // string类型

	// 转uint只能转为uint64
	uint_instance, _ := strconv.ParseUint(numStr, 8, 32) // 将numStr当成2进制 转成32位的uint, 但还想没用~还是字换成64位的uint
	fmt.Printf("uint_instance: %d \n", uint_instance)

	numStr2 := strconv.FormatUint(uint_instance, 2) // 转成2进制的字符串
	fmt.Println("numStr2:", numStr2)
}

type SameFieldA struct {
	name  string
	value int
}

type SameFieldB struct {
	name  string
	value int
}

func (a *SameFieldA) getName() string {
	return a.name
}

func test_struct_convert() {
	fb := SameFieldB{
		name:  "b",
		value: 2,
	}
	fc := SameFieldA(fb) //只要变量名和类型一致就可以进行转换
	n := fc.getName()
	fmt.Println("n:", n)

}

func Run() {
	fmt.Println("test_type_convert")
	//test_int2byte()
	//test_str_convert()
	test_struct_convert()
}
