package model

//定义一个结构体
type student struct {
	Name  string
	score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

func (stu *student) GetScore() float64 {
	return stu.score
}
