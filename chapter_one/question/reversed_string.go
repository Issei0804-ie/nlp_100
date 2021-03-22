package question

import "fmt"

func ReversedString() {
	str := []rune{'s', 't', 'r', 'e', 's', 's', 'e', 'd'}
	result := make(map[int]string)

	for i := 1; i <= len(str); i++ {
		result[i-1] = string(str[len(str)-i])
	}
	fmt.Println()
}
