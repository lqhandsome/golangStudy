package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age int
	Address string
}
var (
	lock sync.Locker
)
func main() {

	var c chan Person
	c = make(chan Person,10)
	Rand(c)
	cLen := len(c)
	for i := 0; i  < cLen;i++ {
		tmp := <-c
		fmt.Println(tmp)
		fmt.Printf("%p",&tmp)
	}
}

func Rand(c chan Person) {
	var person Person
	for i := 1;i <= 10;i++ {
		rand.Seed(time.Now().UnixNano())
		person = Person{
			"a" + strconv.Itoa(i),
			rand.Intn(60),
			"address"  + strconv.Itoa(i),
		}
		c<-person
		fmt.Println(person)
		fmt.Printf("%p",&person)
	}
	fmt.Println("person",person)

}

