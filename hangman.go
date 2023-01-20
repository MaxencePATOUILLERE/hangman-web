package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"math/rand"
	_ "syscall"
	"time"
)

type HangManData struct {
	Save          string
	File          string
	Word          string
	ToFind        string
	Attempts      int
	Used          []rune
	WhichAsciiArt string
	Difficulty    string
	HangManState  []string
}

func setupGameData(ses *sessions.Session) UserData {
	return UserData{
		userName: usernames[len(playerList)],
		userData: HangManData{
			Attempts:     0,
			HangManState: printHangMan(0),
		},
		finish: false,
		uid:    ses.Values["uid"],
		admin:  true,
	}
}

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

func checkInput(data HangManData, l rune) HangManData {
	if l != ' ' && isValid(l) {
		if isUsed(data, l) {
			printHangMan(data.Attempts)
			fmt.Println("Letter already used.")
			return data
		} else if isGood(data.ToFind, l) {
			data = trys(data, l)
		} else {
			data.Attempts++
			data.Used = append(data.Used, l)
			printHangMan(data.Attempts)
			fmt.Println("Not present in the word,", 10-data.Attempts, "attempts remaining")
			return data
		}
	} else {
		fmt.Println("Bad input")
	}
	printHangMan(data.Attempts)

	return data
}

func genMasked(txt string) string {
	result := ""
	for i := 0; i < len(txt); i++ {
		if txt[i] == '_' {
			result = result + "_"
		}
		result = result + "X"
	}
	return result
}

func calcScore(word string) int {
	point := map[rune]int{'a': 1, 'b': 3, 'c': 2, 'd': 2, 'e': 1, 'f': 3, 'g': 3, 'h': 4, 'i': 1, 'j': 4, 'k': 10, 'l': 3, 'm': 2, 'n': 2, 'o': 1, 'p': 2, 'q': 4, 'r': 3, 's': 2, 't': 4, 'u': 1, 'v': 5, 'x': 10, 'y': 10, 'z': 10}
	score := len(word)
	for i := 0; i < len(word); i++ {
		score += point[rune(word[i])]
	}
	return score
}
