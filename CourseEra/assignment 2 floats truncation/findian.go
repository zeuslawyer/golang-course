package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Enter a string: ")

	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("can't read string, reason: %s", err)
	}

	// Here we replace string delimiters and make string lowercase
	// cause we need to case-insensitive
	newLineRemoved := strings.Replace(in, "\n", "", -1)
	newLineRemoved = strings.Replace(newLineRemoved, "\r", "", -1)
	loverCaseInput := strings.ToLower(newLineRemoved)

	startsWithI := strings.Index(loverCaseInput, "i") == 0
	containsA := strings.Contains(loverCaseInput, "a")
	endsWithN := strings.LastIndex(loverCaseInput, "n") == len(loverCaseInput) - 1

	if startsWithI && containsA && endsWithN {
		fmt.Println("Found!")
		return
	}

	fmt.Println("Not Found!")
}