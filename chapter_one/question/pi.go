package question

import (
	"fmt"
	"strings"
)

func Pi() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	str1 := strings.Split(str, " ")
	result := make([]rune, len(str))

	for i := 0; i < len(str1); i++ {
		result[i] = []rune(str1[i])[0]
	}
	fmt.Println(string(result))
}
