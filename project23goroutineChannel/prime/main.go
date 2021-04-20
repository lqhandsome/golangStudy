package main

import "fmt"

func main()  {

	intChan := make(chan int ,8000)
	primeChan := make(chan int ,2000)
	exitChan := make(chan int ,4)
	go putNum(8,intChan)
	for i:=1;i<5;i++ {
		fmt.Println("15")
		go primeNum(intChan,primeChan,exitChan)
	}
	for  {

		if len(exitChan) == 4 {
			break
		}
	}

	fmt.Println("一共有",len(primeChan))

}

func putNum(num int,intChan chan int)  {
	for i := 2;i <= num;i++ {
		intChan<-i
	}
	fmt.Println(len(intChan))
	close(intChan)
}

func primeNum(intChan chan int,primeChan chan  int,exitChan chan  int)  {

	for v := range intChan{
		fmt.Println("读取到了",v)
		count := 0
		for i := 1;i < v;i++ {
			if v % i  == 0{
				count++
			}
		}
		if count <= 1{
			primeChan<-v
		}
	}
	exitChan<-1
}