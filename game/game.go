package game



func Letters() []string {
	var i int = 65
	var answer []string
	for ; i <= 90; i++ {
		answer = append(answer, string(rune(i)))
	}
	return answer
}
