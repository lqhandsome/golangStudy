package main

import "fmt"

func main()  {
	str := "abcd刘强"

	r := []rune(str)
	//按照字符遍历
	for i := 0;i <len(r);i++ {
		//fmt.Println(i,r[i])
		fmt.Printf("%c\n",r[i])

	}
	//按照字符遍历，index还是字节
	for index,val := range str{
		fmt.Println(index,val)
	}
}
