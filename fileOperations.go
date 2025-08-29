package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filePath string) ([]string, error) {
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
			fmt.Println("error reading file: ", err)
			return nil, err
		}

		s = append(s, line)
	}

	return s, nil
}

func saveToFile(filePath string, data []string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("error opening file:", err)
		return
	}
	// Ensure the file is closed after the function completes
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Write a string to the buffer
	for _, line := range data {
		_, err = writer.WriteString(line)
		if err != nil {
			fmt.Println("error writing to buffer:", err)
			return
		}
	}

	// Flush the buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("error flushing buffer:", err)
		return
	}

	fmt.Println("converted SQL is saved to " + filePath)
}
