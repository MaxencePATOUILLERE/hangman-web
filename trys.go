package main

import "fmt"

func trys(data HangManData, testLetter rune) HangManData {
	fmt.Println(string(testLetter))
	if testLetter >= 'A' && testLetter <= 'Z' {
		testLetter += 32
	}
	Index := findIndexLetter(testLetter, data.ToFind)
	addedLetter := []rune(data.Word)
	for i := 0; i < len(data.Word); i++ {
		for j := 0; j < len(Index); j++ {
			if i == Index[j] {
				addedLetter[i] = testLetter
			}
		}
	}
	data.Word = convertInStr(addedLetter)
	data.Used = append(data.Used, testLetter)
	return data
}

func findIndexLetter(TestLetter rune, Words string) []int {
	Index := []int{}
	for i := 0; i < len(Words); i++ {
		if TestLetter == rune(Words[i]) {
			Index = append(Index, i)
		}
	}
	return Index
}

func convertInStr(liste []rune) string {
	finalWord := ""
	for i := 0; i < len(liste); i++ {
		finalWord = finalWord + string(liste[i])
	}
	return finalWord
}

func isGood(str string, test rune) bool {
	for i := 0; i < len(str); i++ {
		if rune(str[i]) == test || rune(str[i]) == test+32 {
			return true
		}
	}
	return false
}

func isUsed(data HangManData, letter rune) bool {
	for i := 0; i < len(data.Used); i++ {
		if data.Used[i] == letter {
			return true
		}
	}
	return false
}
