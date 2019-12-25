// 输入
// i am a student
// 输出
// i ma a tneduts
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str1 string

	fmt.Println("输入一个字符串")

	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	str1, err := inputReader.ReadString('\n')
	if err != nil {
		return
	}
	str_len := len(str1)

	for i, m, n := 0, 0, 0; i < str_len; i, m = i+1, m+1 {
		for ; str1[m] != ' '; m++ {
			if m >= str_len-1 {
				break
			}
		}
		for n = m - 1; n >= i; n-- {
			fmt.Printf("%c", str1[n])
		}
		fmt.Printf("%c", str1[m])
		i = m
	}
}
