package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func getFileWords() string {
	f, err := os.Open("./assets/words/words.txt")
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(f)
	var result []string
	cpt := 0
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
		cpt++
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return result[r.Intn(cpt)]
}

func formatWord(word string) string {
	sSlice := []rune(word)
	for i := 0; i < len(sSlice); i++ {
		if sSlice[i] < 'a' || sSlice[i] > 'z' {
			sSlice[i] = ' '
		}
	}
	return string(sSlice)
}

func genHidden(w string) string {
	hidden := ""
	for i := 0; i < len(w); i++ {
		if w[i] == ' ' {
			hidden += " "
		} else {
			hidden += "_"
		}
	}
	return hidden
}
