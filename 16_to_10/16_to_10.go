// 16_to_10
package main

import (
	"fmt"
)

func main() {
	var num_H string

	fmt.Println("输入16进制数:")
	for {
		_, err := fmt.Scanf("%s", &num_H)
		if err != nil {
			return
		}

		len_H := len(num_H)

		num := 0
		num_D := 0

		for i, j := len_H, 0; i > 2; i, j = i-1, j+1 {
			switch num_H[i-1] {
			case '0':
				num = 0
			case '1':
				num = 1
			case '2':
				num = 2
			case '3':
				num = 3
			case '4':
				num = 4
			case '5':
				num = 5
			case '6':
				num = 6
			case '7':
				num = 7
			case '8':
				num = 8
			case '9':
				num = 9
			case 'A', 'a':
				num = 10
			case 'B', 'b':
				num = 11
			case 'C', 'c':
				num = 12
			case 'D', 'd':
				num = 13
			case 'E', 'e':
				num = 14
			case 'F', 'f':
				num = 15
			}

			for m := 0; m < j; m++ {
				num *= 16
			}
			num_D += num
		}
		fmt.Printf("%s十进制数为%d\n", num_H, num_D)
	}
}
