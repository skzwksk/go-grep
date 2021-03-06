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
	file, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	rep := regexp.MustCompile(os.Args[1])
	for s.Scan() {
		fmt.Println(rep.ReplaceAllStringFunc(s.Text(), getMagentaString))
	}
}

func getMagentaString(s string) string {
	return aurora.Magenta(s).String()
}
