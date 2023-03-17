package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	// 对输入的字符串进行Unicode的检查
	if !utf8.ValidString(s) {
		return s, errors.New("the input value is not Unicode for not valide UTF-8")
	}
	// fmt.Printf("the input: %q \n", s)
	// b := []byte(s) // bug fixed 事实上有用不止一个byte来表示字符的unicode
	b := []rune(s) // bug go在将string转化成rune时 可能会将单字节但是非有效的unicode的string转化成unicode中代表未知字符的字符集 该操作就改变了双反转与原字符串相等的结果
	// fmt.Printf("the rune: %q \n", b)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b), nil
}

func main() {
	input := "the quick brown fox jumped over the lazy dog"

	rev, revErr := Reverse(input)

	doubleRev, doubleRevErr := Reverse(rev)

	fmt.Printf("original: %q \n", input)
	fmt.Printf("reverse: %q, and error: %v\n", rev, revErr)
	fmt.Printf("reverse again: %q, and error: %v\n", doubleRev, doubleRevErr)
}
