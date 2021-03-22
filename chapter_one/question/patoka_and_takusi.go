package question

import "fmt"

func PatokaAndoTakusi() {
	str := "パトカー"
	str1 := "タクシー"
	rune1 := []rune(str)
	rune2 := []rune(str1)
	var result []rune
	for i := 0; i < 4; i++ {
		result = append(result, rune1[i])
		result = append(result, rune2[i])
	}
	fmt.Println(string(result))
}
