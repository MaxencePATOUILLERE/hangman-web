package main

import (
	"math/rand"
	"time"
)

func reveal(data HangManData) HangManData {
	var cpt = 0
	for cpt < len(data.ToFind)/2-1 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		letter := rune(data.ToFind[r.Intn(len(data.ToFind))])
		if !isUsed(data, letter) {
			cpt++
			data = trys(data, letter)
		}
	}
	return data
}

func isValid(l rune) bool {
	if l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z' {
		return true
	}
	return false
}

func finish(data HangManData) bool {
	if data.Word == data.ToFind {
		return true
	} else if data.Attempts == 10 {
		return true
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

func isGood(str string, test rune) bool {
	for i := 0; i < len(str); i++ {
		if rune(str[i]) == test || rune(str[i]) == test+32 {
			return true
		}
	}
	return false
}
