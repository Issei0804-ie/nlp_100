package question

import (
	"fmt"
	"regexp"
	"strings"
)

func NGram() {
	str := "I am an NLPer"
	var str1 [][]string
	splitStr := strings.Split(str, " ")

	fmt.Println("**********************")
	for i := 0; i < len(splitStr); i++ {
		if len(splitStr) < i+2 {
			break
		}
		str1 = append(str1, splitStr[i:i+2])
	}

	for i := 0; i < len(str1); i++ {
		fmt.Println(str1[i])
	}
	fmt.Println("**********************")
	reg := regexp.MustCompile(` `)
	str4 := []rune(reg.ReplaceAllString(str, ""))

	var str5 [][]rune
	for i := 0; i < len(str4); i++ {
		if len(str4) < i+2 {
			break
		}
		str5 = append(str5, str4[i:i+2])
	}
	for i := 0; i < len(str5); i++ {
		fmt.Println(string(str5[i]))
	}
}
