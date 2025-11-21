package test_type

import "fmt"

// u 开头是unsigned 无符号的意思
var a uint8 = 15     // 8个字节的无符号整形 unsigned int (十进制表达)
var b uint8 = 0b1111 // 8个字节的无符号整形 unsigned int (二进制表达)
var c uint8 = 0xF    // 8个字节的无符号整形 unsigned int (十六进制表达)

var d = 15 // 字节长度依赖编译器64位则8个字节(64位)32位则4个字节(32位)

var e float32 = 15
var f float64 = 15
var g complex64 = 10.0 + 5i
var h complex64 = 10.0 - 5i
var r_p = real(h) // 获取复数的实部 10
var i_p = imag(h) // 获取复数的虚部 5
var h_com = complex(10, 5)
var i = "hello"
var s2b = []byte(i)   // string 转 byte
var b2s = string(s2b) // byte 转 string
var j rune = '世'     // ''扩住一个字符 一个字符占用4个字节(32位) 与 int32 类型相同
var k int32 = 1
var s2runes []rune = []rune(i) //string 转rune数组
var m bool = true
var n string = "go世界"
var n2b = []byte(n) // 中文Unicode码点值占用3个字节 ASCII 占用一个字符 byte == uint8 rune==int32
var n2r = []rune(n) // 直接返回每一个字符的Unicode码点

type Gender string

var Male Gender = "male"

func method() (a int8, b string, e string, f string) { // 需要声明返回的数值类型，前面的变量为局部变量可以在方法内使用
	//a := 1 // 返回值里面已经有所定义
	a = 10

	var d, c string = "hello", "world" //多变量赋值
	g, h := 1, "str is here"
	fmt.Println(g)
	fmt.Println(h)
	return a, "hello", d, c
}

func Run() {
	fmt.Println("test_type")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == b) // True
	fmt.Println(a == c) // True
	fmt.Println(c == b) // True
	//fmt.Println(a == d) // 类型不同不能比较
	fmt.Println(e == float32(f)) // 类型不同不能比较,可以强转之后比较
	fmt.Println(g)               // 类型不同不能比较,可以强转之后比较
	fmt.Println(g * h)
	fmt.Println(r_p)
	fmt.Println(i_p)
	fmt.Println(h_com == complex128(h)) // complex64 转 complex128 后比较 有问题?
	fmt.Println(i)
	fmt.Println(s2b)
	fmt.Println(b2s)
	fmt.Println(j == k)
	fmt.Println(j)
	fmt.Println(s2runes)
	fmt.Println(len(s2b))
	fmt.Println(len(i))
	fmt.Println(len(s2runes))
	fmt.Println(len(n))
	fmt.Println(len(n2b))
	fmt.Println(len(n2r))
	fmt.Println("=======")
	fmt.Println(method())

	fmt.Println(Male)
}
