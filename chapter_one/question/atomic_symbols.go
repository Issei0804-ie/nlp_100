package question

import (
	"fmt"
	"strings"
)

func AtomicSymbols() {
	str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	in := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	str2 := strings.Split(str, " ")
	result := make(map[int]string)
	for i, str3 := range str2 {
		if IsContain(in, i+1) {
			result[i] = string([]rune(str3)[0])
		} else {
			result[i] = string([]rune(str3)[0:2])
		}
	}

	for i := 0; i < len(str2); i++ {
		fmt.Println(result[i])
	}

}

func IsContain(in []int, i int) bool {
	for _, a := range in {
		if a == i {
			return true
		}
	}
	return false
}
