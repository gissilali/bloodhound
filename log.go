package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetFileContent(filePath string) ([]string, error) {
	var lines []string
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error while opening file")
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
