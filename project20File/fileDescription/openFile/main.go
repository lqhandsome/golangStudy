package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./description.txt")

	if err != nil {
		fmt.Println("打开失败!",err)
	}
	defer file.Close()
	//带缓冲的读取方式
	reader := bufio.NewReader(file)

	for  {
		str, errs := reader.ReadString('\n')
		if errs == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("读取结束")
}
