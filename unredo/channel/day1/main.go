package main

import "fmt"

func main() {
	c1 := make(chan float64,1)
	c2 := make(chan<- float64,1)
	left(c1)
	c2 = c1
	c2<-1
	fmt.Println(<-c1)
	fmt.Println(c2)
	return
	//go insert(c1)
	//go insert(c1)
	//go insert(c1)
	//go insert(c1)
	//go insert(c1)
	go insert(c1)
	//for  {
	//	val,ok := <-c1 //一直读到退出ok会返回false
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(val)
	//}

	for val := range c1 {
		fmt.Println(val)
	}
}

func insert(c chan float64) {
	i := 1.1
	for {
		c<-i
		i++
		if i == 2.1 {
			close(c)
		}
		select {

		}
	}

}
func left(c chan<- float64) {
	c<-1
}
