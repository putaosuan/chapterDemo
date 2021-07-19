package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

//将Pupil和Graduate共有的方法也绑定到*Student
func (s *Student) ShonInfo() {
	fmt.Printf("姓名=%v,年龄=%v,成绩=%v\n", s.Name, s.Age, s.Score)
}
func (s *Student) SetScore(score float64) {
	s.Score = score
}

//小学生
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中。。。")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试。。。")
}
func main() {
	pupil := &Pupil{}
	pupil.Name = "tom"
	pupil.Age = 18
	pupil.testing()
	pupil.SetScore(70)
	pupil.ShonInfo()

	graduate := &Graduate{}
	graduate.Name = "xiaomi"
	graduate.Age = 10
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShonInfo()
}
