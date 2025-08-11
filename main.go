package main

import (
	"flag"
	"fmt"
	"os"
)

func parseFlags() {
	flag.StringVar(&commandLineArgs.fileName, "f", "", "Define file to read")
	flag.Parse()
}

func main() {
	parseFlags()

	if commandLineArgs.fileName != "" {
		fmt.Println("This is still work in progress: " + commandLineArgs.fileName)
		os.Exit(0)
	}

	args := flag.Args()
	sqlInput := args[len(args)-1]
	output := convertSingleLineToMultilineSQL(sqlInput)

	fmt.Println(output)
}
