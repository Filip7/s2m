package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func parseFlags() {
	flag.StringVar(&cmdLineArgs.fileName, "f", "", "Define file to read")
	flag.StringVar(&cmdLineArgs.outputFileName, "o", "", "Define file to save the output to")
	flag.Parse()
}

func main() {
	parseFlags()

	if cmdLineArgs.fileName != "" {
		out, err := readFile(cmdLineArgs.fileName)
		if err != nil {
			fmt.Println("error happened", err)
			os.Exit(1)
		}

		converted := ConvertSingleLineToMultilineSQLFromFile(out)
		if cmdLineArgs.outputFileName != "" {
			saveToFile(cmdLineArgs.outputFileName, converted)
		} else {
			fmt.Println(strings.Join(converted, ""))
		}
		os.Exit(0)
	}

	args := flag.Args()
	sqlInput := args[len(args)-1]
	converted := ConvertSingleLineToMultilineSQL(sqlInput)

	if cmdLineArgs.outputFileName != "" {
		saveToFile(cmdLineArgs.outputFileName, converted)
	} else {
		output := strings.Join(converted, "")
		fmt.Println(output)
	}
}
