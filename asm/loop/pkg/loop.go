package pkg


func LoopAdd(cnt, v0, step int) int {
	result := v0
	for i := 0; i< cnt; i++ {
		result += step
	}
	return result
}

func AsmLoopAdd(cnt, v0, step int) int