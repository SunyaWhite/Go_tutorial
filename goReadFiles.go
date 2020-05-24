package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkLastLine(err error) bool {
	if err != nil && err.Error() == "EOF" {
		return true
	}
	return false
}

func readLinesFromFile(absolutePath *string) []string {
	file, err := os.Open(*absolutePath)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	result := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if checkLastLine(err) {
			break
		}
		result = append(result, string(line))
	}
	fmt.Println("Reading is done")
	return result
}

func writeToFile(absolutePath *string, lines *[]string) {
	file, err := os.Create(*absolutePath)
	check(err)
	defer file.Close()
	for _, value := range *lines {
		_, err := file.WriteString("line " + value + "\n")
		check(err)
	}
	fmt.Println("Writing is done")
}

func displayLines(lines *[]string) {
	for index, value := range *lines {
		fmt.Println("line " + string(index) + " " + value)
	}
}

func main() {
	fmt.Printf("HelloWorld\n")
	pathToRead := "/home/sunya/Documents/TempFiles/LogFile.txt"
	pathToWrite := "/home/sunya/Documents/TempFiles/ExampleFile.txt"
	lines := readLinesFromFile(&pathToRead)
	//displayLines(&lines)
	writeToFile(&pathToWrite, &lines)
}
