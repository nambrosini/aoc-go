package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(day int) string {
	path := fmt.Sprintf("inputs/input%d.txt", day)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Couldn't close file", err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var fileContent string
	for scanner.Scan() {
		line := scanner.Text()
		fileContent += line + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file", err)
	}

	return fileContent[:len(fileContent)-1]
}
