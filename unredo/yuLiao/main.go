package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	writeFile := "图纸名.txt"
	file, err := os.OpenFile(writeFile, os.O_APPEND|os.O_CREATE, 777)
	defer file.Close()
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	for i := range fileInfoList {
		arr := strings.Split(fileInfoList[i].Name(), "-")
		for _ , val := range arr {
				file.Write([]byte(string(val)))
				file.Write([]byte{'\t'})

		}
		file.Write([]byte{'\n'})
	}
}
