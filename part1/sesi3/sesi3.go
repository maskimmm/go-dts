package main

import "fmt"

func main() {
	calculateWord("selamat malam")
}

func calculateWord(words string) {
	var wordCount map[string]int
	wordCount = map[string]int{}
	wordLength := len([]rune(words))
	for x := 0; x < wordLength; x++ {
		value, exist := wordCount[string(words[x])]
		_ = value
		if exist {
			wordCount[string(words[x])] += 1
		} else {
			wordCount[string(words[x])] = 1
		}
		fmt.Printf("%c\n", words[x])
	}
	fmt.Println(wordCount)
}
