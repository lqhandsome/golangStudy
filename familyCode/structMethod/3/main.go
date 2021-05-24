package main

import "fmt"



func main()  {
	person := new(Person)
	person.name = "LQ"
	person.updateName("hello world")
	//fmt.Println(person)

	var p Person
	p.name = "1"
	//fmt.Printf("%p,%T,%v\n",person,person,person)
	fmt.Println(&p)
	return
	person_new := person
	person_new.name = "mew"
	fmt.Println(person)
	return

	fmt.Println(person.updateName("handsome"))
	p1 := Person{
		"lq",
		23,
		":",
	}
	deleteName(person)
	fmt.Println(p1)
	fmt.Println(person)
}

type Person struct {
	name string
	Age int
	Like string
}

func (p *Person) updateName(newName string)(string,error){
	p.name =  newName +p.name
	//fmt.Printf("%p,%T %v\n ",&p,p,p)
	return newName +p.name ,nil
}

func deleteName(p *Person){
	p.name = "11"
}

func (p Person) String()string{
	//str := fmt.Sprintf("%v")
	fmt.Println("hello world")
	return "this is string"
}