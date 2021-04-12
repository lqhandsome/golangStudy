package main

import (
	"fmt"
	"math/rand"
	"sort"
	//"time"
)
//声明结构体
type Hero struct {
	name string
	age int
}

//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)

//申明类型
type HeroSlice []Hero

func (he HeroSlice) Len()int{
	return len(he)
}

func (he HeroSlice) Less(i, j int)bool{
	return he[i].age < he[j].age
}

func (he HeroSlice) Swap(i, j int){
	he[i],he[j] = he[j],he[i]
}


func main() {
	rand.Seed(80)
	var heroes HeroSlice
	for i := 0;i < 10;i++ {
		hero := Hero{
			fmt.Sprintf("英雄代号%d",rand.Intn(1000)),
			rand.Intn(1000),
		}

		heroes = append(heroes,hero)
	}

	sort.Sort(heroes)
	fmt.Println(heroes)

}
