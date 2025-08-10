package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func convertSingleLineToMultilineSQL(input string) string {
	lines := strings.Split(input, ";")
	lines = removeEmptyStrings(lines)

	// (?i) -> case insensitive
	// w+   -> match word
	// (.*) -> match everything inside '(' and ')'
	r := regexp.MustCompile(`(?i)insert into \w+(.*) values `)
	match := r.FindAllString(lines[0], -1)
	var match1 string
	if match == nil {
		match1 = ""
	} else {
		match1 = match[0]
	}

	isFirst := 0
	for i := range len(lines) {
		match = r.FindAllString(lines[i], -1)
		var match2 string
		if match != nil {
			match2 = match[0]
		} else {
			match2 = ""
		}

		if match1 != match2 {
			isFirst = 0
		}

		if i != len(lines)-1 {
			lines[i] += ","
		} else {
			lines[i] += ";"
		}

		// if new insert block, replace prev ',' with ';'
		if isFirst == 0 && i != 0 {
			lines[i-1] = strings.TrimRight(lines[i-1], ",")
			lines[i-1] += ";"
		}

		if isFirst != 0 {
			// remove the insert part and replace it with nothing
			lines[i] = strings.Replace(lines[i], match2, "", 1)
			// remove the starting newline char that we got at the begining by spliting at ;
			lines[i] = strings.Replace(lines[i], "\n", "", 1)
			// return newline char and add spaces for the len of insert part, to pretify the output
			lines[i] = "\n" + strings.Repeat(" ", len(match2)) + lines[i]
		}
		match1 = match2
		isFirst += 1
	}

	return strings.Join(lines, "")
}

func main() {
	flag.Parse()

	args := flag.Args()
	sqlInput := args[len(args)-1]
	output := convertSingleLineToMultilineSQL(sqlInput)

	fmt.Println(output)
}
