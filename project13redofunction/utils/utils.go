package utils

import "fmt"
var I = "utils"
func init() {
	I = "lq"
	fmt.Println("utils.init")
}

func main()  {
	fmt.Println("utils.main")
}