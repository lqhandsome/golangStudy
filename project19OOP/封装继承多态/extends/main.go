package main

import "fmt"

type Student struct {
	name string
	age int
	score float64
}

func (stu *Student) SetScore(score float64){
	stu.score = score
}

func (stu *Student) ShowScore(){
	fmt.Println(stu.score)
}
type Pupil struct {
		Student
		love string

}
func (stu *Pupil) showLove(){
	fmt.Println(stu.love)
}
func main() {
	ss := Pupil{}
	ss.SetScore(500)
	ss.name = "l1"
	ss.love = "cpu"
	ss.Student.ShowScore()
	fmt.Println(ss)
}

//单例模式
func AddUpper(n int) func(int)int{
	return func (x int) int{
		n = n + x
		return n
	}
}