package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filePath string) (string, error) {
	var s []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading file: ", err)
			return "", err
		}

		s = append(s, line)
	}

	out := ConvertSingleLineToMultilineSQLFromFile(s)
	return out, nil
}
