package question

import "fmt"

func Patatotakasi() {
	str := "パタトクカシーー"
	runeStr := []rune(str)
	var result = make([]rune, 4)
	for i := 0; i <= 3; i++ {
		result[i] = runeStr[i*2+1]
	}
	fmt.Println(string(result))
}
