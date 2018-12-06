package main

import (
	"bufio"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"strings"
)

func main() {
	fileText := getFileText(os.Args[2])
	searchWord := os.Args[1]
	index := strings.Index(fileText, searchWord)
	if index >= 0 {
		fmt.Println(strings.Replace(fileText, searchWord, aurora.Magenta(searchWord).String(), -1))
	}
}

func getFileText(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var fileText string
	for s.Scan() {
		fileText = fileText + s.Text()
	}
	return fileText
}
