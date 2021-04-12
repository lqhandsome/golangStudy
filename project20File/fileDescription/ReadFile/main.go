package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "./file.txt"
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	fmt.Printf("%c",(content))
}
