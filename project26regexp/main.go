package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println(regexp.Match("^(He).* ", []byte("Hello World!")))
}
