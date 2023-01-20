package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func save(data HangManData) bool {
	var name string
	fmt.Print("Choose a valid filename to save : ")
	fmt.Scan(&name)
	if !isAlreadyExist(name) {
		var choice string
		fmt.Print("File already exist overwrite it ? (Y/N) ")
		fmt.Scan(&choice)
		if choice == "Y" || choice == "y" || choice == "Yes" || choice == "yes" {
			os.Remove(name + ".txt")
		} else {
			return false
		}
	}
	content, _ := json.Marshal(data)
	_ = os.WriteFile(name+".txt", content, 0644)
	fmt.Println("Game Saved in " + name + ".txt")
	return true
}

func isAlreadyExist(path string) bool {
	_, err := os.Stat(path + ".txt")
	if err != nil {
		return true
	}
	return false
}

func saveWithPath(data HangManData, path string) {
	content, _ := json.MarshalIndent(data, "", " ")
	if isAlreadyExist(path) || isAlreadyExist(path+".txt") {
		_ = os.Remove(path)
		_ = os.WriteFile(path, content, 0644)
	} else {
		_ = os.Remove(path + ".txt")
		_ = os.WriteFile(path+".txt", content, 0644)
	}
	fmt.Println("Game Saved in " + path)
	return
}

func getFileData(file *string) HangManData {
	var GameData HangManData
	jsonFile, err := os.Open(*file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &GameData)
	if !isValidJSON(GameData) {
		return GameData
	}
	fmt.Print("\033[H\033[2J")
	printStart()
	fmt.Println("Welcome Back, you have " + string(rune(GameData.Attempts)+48) + " attempts remaining.")
	GameData.Save = *file
	return GameData
}

func isFileValid(file string) bool {
	fileContent, _ := os.Open(file)
	scanner := bufio.NewScanner(fileContent)
	var result []string
	cpt := 0
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
		cpt++
	}
	if !isAlreadyExist(file) {
		return false
	}
	if !isValidJSON(getFileData(&file)) {
		return false
	}
	return true
}

func isValidJSON(GameData HangManData) bool {
	if GameData.Word == "" || GameData.ToFind == "" || len(GameData.Used) == 0 {
		return false
	}
	return true
}
