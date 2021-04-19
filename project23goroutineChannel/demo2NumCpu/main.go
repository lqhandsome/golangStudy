package main

import (
	"fmt"
	"runtime"
)
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
	fmt.Println(runtime.GOMAXPROCS(cpuNum-1))
	fmt.Println(runtime.NumCPU())
}
