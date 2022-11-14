package main

import (
	"fmt"
	"strings"
)

func guessWord(data HangManData, letter string) HangManData {
	var response string
	fmt.Print("You have entered more than one letter want to try to guess the word ? (Y/N)")
	fmt.Scan(&response)
	letter = strings.ToLower(letter)
	if response == "Y" || response == "Yes" || response == "y" || response == "yes" {
		if letter == data.ToFind {
			data.Word = data.ToFind
			return data
		} else {
			data.Attempts += 2
			printHangMan(data.Attempts)
			fmt.Println("Bad word try again ! " + string(rune(-data.Attempts+58)) + " left.")
		}
	} else if response != "N" && response != "No" && response != "n" && response != "no" {
		printHangMan(data.Attempts)
		fmt.Println("Bad input try again")
	}
	printWord(data)
	return data
}
