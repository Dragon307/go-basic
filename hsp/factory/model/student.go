package model

// 定义一个结构体
type student struct {
	Name string
	score float64
}
// 1, 首字母小写的结构体只能在当前使用， 2， 字段小写属于私有

// 因为 student 结构体首字母是小写，因此是只能在model使用
// 我们使用工厂模式来解决

func NewStudent(n string, s float64) *student {
	return &student{
		Name: n,
		score: s,
	}
}

// 如果 score 字段首字母小写， 则，在其它包不可以直接使用，我们可以提供一个方法
func (s *student) GetScore() float64 {
	return s.score
}