package question

import "fmt"

func ReversedString() {
	str := []rune{'s', 't', 'r', 'e', 's', 's', 'e', 'd'}
	result := make([]rune, len(str))
	for i := 1; i <= len(str); i++ {
		result[i-1] = str[len(str)-i]
	}
	fmt.Println(string(result))
}
