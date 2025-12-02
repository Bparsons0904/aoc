package utilities

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) []string {
	file, err := os.Open(fmt.Sprintf("./files/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fileLines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fileLines
}
