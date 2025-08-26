package main

import (
	"flag"
	"fmt"
	"os"
)

func parseFlags() {
	flag.StringVar(&cmdLineArgs.fileName, "f", "", "Define file to read")
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

		fmt.Println(out)
		os.Exit(0)
	}

	args := flag.Args()
	sqlInput := args[len(args)-1]
	output := ConvertSingleLineToMultilineSQL(sqlInput)

	fmt.Println(output)
}
