package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	fileName := "../description.txt"

	file, err := os.OpenFile(fileName,os.O_RDWR | os.O_APPEND,7)

	if err != nil {
		fmt.Println("打开失败")
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for  {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
	file.WriteString("hello\r\n")
}
