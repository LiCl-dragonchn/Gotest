package main

import "fmt"

func main()  {
	var str1 string
	var num int

	_,err := fmt.Scanln(&num)
	if(err != nil){
		return
	}
	_,err = fmt.Scanln(&str1)
	if(err != nil){
		return
	}

	str2 := str1[:num]
	str3 := str1[num:]
	fmt.Println(str3+str2)
}
