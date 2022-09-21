package utils

import (
	"regexp"
)

func is_all_chinese(input string) bool {

	var a = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	for _, v := range input {
		if a.MatchString(string(v)) {
			return true
		}
	}
	return false

}
