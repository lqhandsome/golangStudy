package main

import (
	"fmt"
	"os"
)

//带缓冲的读取方式
//func main() {
//	fileName := "test.txt"
//	res ,err := os.Open(fileName)
//	if err != nil {
//		fmt.Println("打开失败",err)
//	}
//	reader := bufio.NewReader(res)
//
//
//	for  {
//		line ,err:= reader.ReadString('\n')
//		if err ==  io.EOF {
//			fmt.Println(strings.Trim(line,"\r\n"))
//			fmt.Println("读取完成")
//			return
//		}
//		fmt.Println(strings.Trim(line,"\r\n"))
//	}
//}
//直接读取
func main() {
	fileName := "test.txt"
	//方式1
	/*
	res ,_ := os.Open(fileName)
	 buf := make([]byte,1024)
	n,_ := res.Read(buf)
	fmt.Println(string(buf[:n]))
	*/
	//方式2
	//content,_ := ioutil.ReadFile(fileName)
	//fmt.Println(string(content))

	//写入文件
	res,err := os.OpenFile(fileName,os.O_APPEND|os.O_CREATE,7)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := "hello"
	res.Write([]byte(str))
}