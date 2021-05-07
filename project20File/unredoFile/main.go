package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "a.txt"
	//file, err := os.OpenFile(fileName,os.O_APPEND | os.O_RDWR | os.O_RDONLY,0777)
	//if err != nil {
	//	fmt.Println("打开文件失败",err)
	//	return
	//}
	//var tempByte = make([]byte, 128)
	//n,err := file.Read(tempByte)
	//if err != nil {
	//	fmt.Println("打开文件失败",err)
	//	return
	//}
	//fmt.Println(string(tempByte))
	//fmt.Println("读取到的字节数",n)
	//file.Close()
	////err = os.Remove(fileName)
	//if err != nil {
	//	fmt.Println("打开文件失败",err)
	//	return
	//}
	str := "hello/hello"
	err := ioutil.WriteFile(fileName,[]byte(str),777)
	if err != nil {
		fmt.Println(err)
	}
}
