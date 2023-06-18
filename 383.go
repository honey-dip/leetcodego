package leetcodego

import "strings"

func canConstruct(ransomNote string, magazine string) bool {
	atoz := "abcdefghijklmnopqrstuvwxyz"
	result := true
	for _, c := range atoz {
		if strings.Count(ransomNote, string(c)) > strings.Count(magazine, string(c)) {
			result = false
			break
		}
	}
	return result

}
