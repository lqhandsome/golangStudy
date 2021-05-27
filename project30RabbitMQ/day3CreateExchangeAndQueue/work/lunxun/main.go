package main

import "fmt"

func main()  {
	fmt.Println("f1 result: ", f1())
	fmt.Println("f2 result: ", f2())

}

func f1() int {
	var i int
	defer func() {
		i++
		fmt.Println("f11: ", i)
	}()

	defer func() {
		i++
		fmt.Println("f12: ", i)
	}()

	i = 1000
	return i
}

func f2() (i int) {
	defer func() {
		i++
		fmt.Println("f21: ", i)
	}()

	defer func() {
		i++
		fmt.Println("f22: ", i)
	}()

	i = 1000
	return i
}