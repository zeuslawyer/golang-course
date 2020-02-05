package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fName string
	lName string
}

var persons []Person

const MaxFieldSize = 20

func main() {

	var fileName string
	fmt.Printf("Enter File Name (for example: ./test.txt): ")
	if _, err := fmt.Scan(&fileName); err != nil {
		log.Fatalf("Can't read fileName, reason: %v\n", err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Can't open file '%s', reason: %v\n", fileName, err)
	}

	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')

		strSlice := strings.Split(strings.Replace(str, "\n", "", -1), " ")

		var fName, lName string
		if len(strSlice) > 1 {
			fName = strSlice[0]
			lName = strSlice[1]
		} else if len(strSlice) == 1 {
			fName = strSlice[0]
			lName = ""
		}

		if len(fName) > MaxFieldSize {
			fName = fName[0:MaxFieldSize]
		}

		if len(lName) > MaxFieldSize {
			lName = lName[0:MaxFieldSize]
		}

		persons = append(persons, Person{
			fName: fName,
			lName: lName,
		})

		if err != nil {
			break
		}
	}

	if err = f.Close(); err != nil {
		log.Fatalf("Can't close file '%s', reason: %v\n", fileName, err)
	}

	if len(persons) > 0 {
		for i, v := range persons {
			fmt.Printf("%d) FirstName: '%s', LastName: '%s'\n", i+1, v.fName, v.lName)
		}
	}
}