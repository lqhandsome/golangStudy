package main

import (
	"fmt"
	"go_code/project19OOP/factory/model"
)
func main() {
	var stu = model.GetStuden("lq",98.9)

	fmt.Println(*stu)
}
