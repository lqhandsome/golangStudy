package main


import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)
func main() {
	fileName := "../ReadFile.file.txt"
	file, err := os.OpenFile(fileName,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("open file error=%v\n",err)
		return
	}
	defer file.Close()
	str := "hello world\n"
	writer := bufio.NewWriter(file)
	for i := 0;i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	content,_ := ioutil.ReadFile(fileName)
	fmt.Println(content)
}
