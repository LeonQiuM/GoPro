package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 字符串和字节数组
/*
bytes： API 都只针对字节数组
strings: 对字符串的查询、替换、比较、截断、拆分和合并

strconv： 布尔型、整型、浮点型等直接的装换
unicode： 判断字符串的类型，大小写等
*/
func main() {
	fmt.Println(basename("a/dn/c/goo"))
	fmt.Println(basename("a/dn/c/goo.go.1111"))
	fmt.Println(strings.Contains("Hello", "H"))        //包含 返回 bool
	fmt.Println(strings.Count("asdasdqwevergre", "s")) //  统计字符串出现的次数 返回整数
	fmt.Println(strings.HasPrefix("Hello", "He"))      // 前缀
	fmt.Println(strings.HasSuffix("Hello", "lo"))      // 后缀
	fmt.Println(strings.Index("Hello", "lo"))          // 从前往后查，返回索引，没查到是-1
	var ss = []string{"aa", "bb", "cc"}
	fmt.Println(strings.Join(ss, "--")) // 对字符串数组按照给定分隔符进行拼接

	// bytes

	s := "Hello" // 字符串只读，但是需要修改的话就可以将其转换为字节数组，然后进行修改
	b := []byte(s)
	fmt.Println(b)
	b[1] = 'x'
	s2 := string(b)
	fmt.Println(s2)

	var abc = []int{1, 2, 3, 4, 5}
	fmt.Println(intToString(abc))

}

func basename(path string) string {
	index := strings.LastIndex(path, "/") // 从后面开始查找，得到切片后的索引号
	fmt.Println(index)
	s := path[index+1:]
	fmt.Println(s)
	//dot := strings.LastIndex(s, ".")
	//s = s[:dot]
	//
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		fmt.Println(dot)
		s = s[:dot]
	} else {
		s = "err"
	}
	return s

}

func intToString(value []int) string {
	/*
		将 int 类型的数组转换为字符串类型
	*/
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range value {
		if i > 0 {
			buf.WriteByte(',')
		}

		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
