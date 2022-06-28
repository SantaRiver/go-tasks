package anagram

import (
	"sort"
	"strings"
)

func FormatString(word string) (sortWord string) {
	word = strings.Join(strings.Fields(word), "")
	word = strings.Replace(word, ",", "", -1)
	word = strings.Replace(word, "?", "", -1)
	word = strings.Replace(word, "!", "", -1)
	word = strings.Replace(word, string('"'), "", -1)
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, "'", "", -1)
	word = strings.Replace(word, ":", "", -1)
	word = strings.Replace(word, "(", "", -1)
	word = strings.Replace(word, ")", "", -1)
	word = strings.ToLower(word)
	wordR := []rune(word)
	sort.Slice(wordR, func(i int, j int) bool { return wordR[i] < wordR[j] })
	sortWord = string(wordR)
	return sortWord
}

func FindAnagrams(dictionary []string, word string) (result []string) {
	if word != " " && word != "" {
		for _, dWord := range dictionary {
			if FormatString(dWord) == FormatString(word) && strings.ToLower(dWord) != strings.ToLower(word) {
				result = append(result, dWord)
			}
		}
	}
	return result
}
