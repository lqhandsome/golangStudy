package main


import (
	"os"
	"fmt"
)
func main() {
	fmt.Println("命令行的参数",len(os.Args))
	for i,v := range os.Args {
		fmt.Println(i,v)
	}
	/*
	0 main.exe
	1 tmo
	2 ./config.ini
	3 99
	*/
}
