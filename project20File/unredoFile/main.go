package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "./a.txt"
	file, err := os.OpenFile(fileName,os.O_APPEND | os.O_RDWR,777)
	if err != nil {
		fmt.Println("打开失败err=",err)
		return
	}
	reader := bufio.NewReader(file)
	for  {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	file.WriteString("hello world")
	defer file.Close()
	str ,_:= ioutil.ReadFile(fileName)
	fmt.Println(string(str))
}
