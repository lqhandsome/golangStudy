package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
func main() {
	pwd, _ := os.Getwd()
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	writeFile := "图纸名.txt"
	file ,err := os.OpenFile(writeFile,os.O_APPEND|os.O_CREATE,777)
	defer file.Close()
	if err != nil {
		fmt.Println("err=",err)
		return
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		file.Write([]byte(fileInfoList[i].Name()))
		file.Write([]byte{'\n'})
	}
}