package main

import (
	"bufio"
	"fmt"
	"os"
)

func printHangMan(failAttempts int) []string {
	res := []string{}
	hangMan := openHangManTxt()
	if failAttempts == 0 {
		for k := 0; k < 8; k++ {
			res = append(res, hangMan[0])
		}
	} else {
		for i := 0; i < 8; i++ {
			res = append(res, hangMan[i+8*(failAttempts-1)])
		}
	}
	return res
}

func printWord(data HangManData) {
	if data.WhichAsciiArt != "" {
		printASCIIArt(data)
	} else {
		for i := 0; i < len(data.Word); i++ {
			fmt.Print(string(data.Word[i]))
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func printStart() {
	fmt.Println("------------------------------")
	fmt.Println("      Welcome to Hangman      ")
	fmt.Println("------------------------------")
}

func openHangManTxt() []string {
	var f *os.File
	f, _ = os.Open("./assets/hangman.txt")
	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
