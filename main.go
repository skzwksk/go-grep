package main

import (
	"bufio"
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
	"os"
	"regexp"
)

func main() {
	fileText := getFileText(os.Args[2])
	searchWord := os.Args[1]
	rep := regexp.MustCompile(searchWord)
	fmt.Println(rep.ReplaceAllStringFunc(fileText, getMagentaString))
}

func getFileText(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var fileText string
	for s.Scan() {
		fileText = fileText + s.Text()
	}
	return fileText
}

func getMagentaString(s string) string {
	return aurora.Magenta(s).String()
}
