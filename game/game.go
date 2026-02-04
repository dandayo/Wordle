package game

const word1 string = "hello"
const word2 string = "admin"
const word3 string = "amigo"

func Letters() string {
	var i int = 65
	var answer string
	for ; i <= 90; i++ {
		answer += string(rune(i))
	}
	return answer
}

func CheckLetters(input string) string {
	var check string
	for _, v := range input {
		check += string(v - 'a' + 'A')

	}
	return check
}

func UpdateLetters(word string) string {
	old := Letters()
	new := CheckLetters(word)
	var newAnswer string

	for i := 0; i < len(old); i++ {
		for _, v := range new {
			if rune(old[i]) == v {
				i++
			}
		}
		newAnswer += string(rune(old[i]))
	}
	return newAnswer
}
