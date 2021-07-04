package main

import (
	"flag"
	"fmt"
	"os"
)
var name string
func init() {
	flag.StringVar(&name,"name","everyone","to one?")
	//也可以使用 var name = flat.String("name","everyone","the detail")
	//该包的说明
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,"Usage of %s:\n","Question")
		flag.PrintDefaults()
	}
}
func main() {
	//解析参数
	flag.Parse()
	fmt.Printf("Hello,%s\n",name)
}
