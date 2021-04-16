package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	readFileName := "./file.txt"
	writeFileName := "./writeFile.txt"
	content, _ := ioutil.ReadFile(readFileName)
	erra := ioutil.WriteFile(writeFileName,content,0666)
	if erra != nil {
		fmt.Println("写文件出错了")
		return
	}
	fmt.Println(reflect.TypeOf(erra))
	fmt.Println(string(content))
}
