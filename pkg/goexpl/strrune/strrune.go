// Package strrune
/*
	这个package 用于 string 和 rune 类型的 example
*/
package strrune

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// StringRune
/*
	golang 的 string 是一个只读（read-only）的 byte 类型的 slice，即 []byte
	该语言和标准库特别对待字符串 - 作为以 UTF-8 编码的文本的容器。
	在 golang 中，字符被称为“rune”，它是一个代表 utf-8 字符的整数
*/
type StringRune struct {
}

func (s *StringRune) RunExample(inputParams *goexpl.InputParams) error {
	const str1 = "hello"
	const str2 = "哈啰"

	fmt.Println("Len:", len(str1))
	fmt.Println("Num of Rune:", len([]rune(str1)))
	fmt.Println(str1)
	fmt.Printf("存在于 ascii 中的字符，每个字符 %d byte\n", len(str1)/len([]rune(str1)))
	fmt.Println("Len:", len(str2))
	fmt.Println("Num of Rune:", len([]rune(str2)))
	fmt.Println(str2)
	fmt.Printf("存在于 ascii 中的字符，每个字符 %d byte\n", len(str2)/len([]rune(str2)))

	byteSlc1 := []byte(str1)
	fmt.Println("byte slice 长度：", len(byteSlc1))
	fmt.Println("byte slice 内容：", byteSlc1)
	fmt.Println("string 转 byte slice：", []byte(str1))
	fmt.Println("string 转 rune slice：", []rune(str1))

	byteSlc2 := []byte(str2)
	fmt.Println("byte slice 长度：", len(byteSlc2))
	fmt.Println("byte slice 内容：", byteSlc2)
	fmt.Println("string 转 byte slice：", []byte(str2))
	fmt.Println("string 转 rune slice：", []rune(str2))

	return nil
}

func (s *StringRune) Init() {
	if err := goexpl.RegisterGoExample(s.GetGoExampleName(), &StringRune{}); err != nil {
		panic(err)
	}
}

func (s *StringRune) GetGoExampleName() string {
	return common.StringRune
}
