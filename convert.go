package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, strings.TrimLeftFunc(str, unicode.IsSpace))
		}
	}
	return r
}

func replaceLastRune(s string, new rune) string {
	return s[:len(s)-2] + string(new)
}

func addCorrectEnding(i int, lines []string, isFirst int, fromFile bool) {
	if fromFile {
		if i != len(lines)-1 {
			lines[i] = replaceLastRune(lines[i], ',')
		} else {
			lines[i] = replaceLastRune(lines[i], ';')
		}

		// if new insert block, replace prev ',' with ';'
		if isFirst == 0 && i != 0 {
			lines[i-1] = strings.TrimRight(lines[i-1], ",")
			lines[i-1] += ";\n"
		}
	} else {
		if i != len(lines)-1 {
			lines[i] += ","
		} else {
			lines[i] += ";\n"
		}

		// if new insert block, replace prev ',' with ';'
		if isFirst == 0 && i != 0 {
			lines[i-1] = strings.TrimRight(lines[i-1], ",")
			lines[i-1] += ";\n"
		}
	}
}

func ConvertSingleLineToMultilineSQL(input string) []string {
	lines := strings.Split(input, ";")
	lines = removeEmptyStrings(lines)

	convert(lines, false)

	return lines
}

func ConvertSingleLineToMultilineSQLFromFile(input []string) []string {
	input = removeEmptyStrings(input)

	convert(input, true)

	return input
}

func convert(lines []string, fromFile bool) {
	// (?i)  -> case insensitive
	// (\w+) -> match word
	// (.*)  -> match everything inside '(' and ')'
	r := regexp.MustCompile(`(?i)insert into (\w+)(.*) values `)

	isFirst := 0
	match, table := "", ""
	for i := range len(lines) {
		firstWord := fmt.Sprintf("%.7s", lines[i])
		if strings.Contains(strings.ToLower(firstWord), "insert") {
			// is an insert
			if isFirst == 0 {
				insertMatch := r.FindStringSubmatch(lines[i])
				match = insertMatch[0]
				table = insertMatch[1]
			}
			// get first 4 words, and get the table name from it
			splitLine := strings.SplitN(lines[i], " ", 5)
			locatedTableName := splitLine[len(splitLine)-3]
			if locatedTableName != table {
				// if new table insert has been found
				isFirst = 0
				addCorrectEnding(i, lines, isFirst, fromFile)
				insertMatch := r.FindStringSubmatch(lines[i])
				match = insertMatch[0]
				table = locatedTableName
				isFirst = 1
				continue
			}
		} else {
			// is not an insert
			isFirst = 0
			addCorrectEnding(i, lines, isFirst, fromFile)
			continue
		}

		addCorrectEnding(i, lines, isFirst, fromFile)

		if isFirst != 0 {
			// remove the insert part and replace it with nothing
			lines[i] = strings.Replace(lines[i], match, "", 1)
			// remove the starting newline char that we got at the begining by spliting at ;
			lines[i] = strings.Replace(lines[i], "\n", "", 1)
			// return newline char and add spaces for the len of insert part, to pretify the output
			lines[i] = "\n" + strings.Repeat(" ", len(match)) + lines[i]
		}
		isFirst += 1
	}
}
