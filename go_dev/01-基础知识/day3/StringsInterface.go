package main

import (
	"fmt"
	"strings"
)

func UrlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("https://%s", url)
	}
	return url

}
func PathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path

}

func main() {
	var (
		url string
	)
	fmt.Scanf("%s", &url)
	url = UrlProcess(url)
	fmt.Println(url)
	url = PathProcess(url)
	fmt.Println(url)

	res := strings.Index("abcedasdasde", "sd") //  从前向后检查，返回位置，没有返回-1
	fmt.Printf("res:%d\n", res)
	res1 := strings.LastIndex("abcedasdasde", "sd") //  从后向前检查，返回位置，没有返回-1
	fmt.Printf("res:%d\n", res1)

	/*
		strings.Replace(str string, old string, new string, n int)string //  在源字符串中将新字符串替换到老字符串的位置，并设定替换多少次， 全部替换为-1， 超过不会报错， o次不替换
	*/
	var str1 string = "abcdefgab"
	res2 := strings.Replace(str1, "ab", "AB", 0)
	fmt.Println(res2)

	/*
		strings.Count(str string,substr string) int 字符串计数
	*/

	/*
		strings. Repeat(str string, count int) string  将多个 str重复cout 次拼接起来
	*/
	res3 := strings.Repeat(str1, 3)
	fmt.Println(res3)

	/*
		strings.ToLower  转小写，字符串中已经是大写的不会报错
		strings.ToUpper  转大写，字符串中已经是小写的不会报错

	*/

	fmt.Println(strings.ToLower("ABc"))
	fmt.Println(strings.ToUpper("Abc"))

	//  去掉字符串首位的空包字符
	fmt.Println(len(strings.TrimSpace(" asdbac ")))

	// 去掉字符串首尾固定字符,挨着有多少个[连续]去除多少个
	fmt.Println(strings.Trim("AABBCCAAAAA", "A"))

	// strings.TrimLeft
	fmt.Println(strings.TrimLeft("AABBCCAAAAA", "A"))

	// strings.TrimRight
	fmt.Println(strings.TrimRight("AABBCCAAAAA", "A"))

	// 按照空格拆分成 slice
	fmt.Println(strings.Fields("AB a asd 3e"))

	// 指定分隔符返回 slice
	fmt.Println(strings.Split("AB/a/asd/3e", "/"))

	// 指定拼接符号拼接
	var str22 = []string{"ad", "ba", "ca"}
	fmt.Println(strings.Join(str22, "="))

}
