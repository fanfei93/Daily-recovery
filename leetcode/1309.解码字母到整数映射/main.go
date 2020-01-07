package main

import (
	"strconv"
)

func freqAlphabets(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) {
			if s[i+2] == '#' {
				strInt, _ := strconv.Atoi(string(s[i : i+2]))
				res += string(byte(strInt + 96))
				i = i + 2
				continue
			}
		}
		res += string(s[i] + 48)
	}
	return res
}
