package main

import "strings"

func main() {

}

func findFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	if len(str) == 0 || indexFirstBracketFound < 0 {
		return ""
	}

	runes := []rune(str)
	wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound >= 0 {
		runes := []rune(wordsAfterFirstBracket)
		return string(runes[1 : indexClosingBracketFound-1])
	}

	return ""
}
