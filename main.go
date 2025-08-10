package main

import (
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

	for i := range len(lines) {
		if i != len(lines)-1 {
			lines[i] += ","
		} else {
			lines[i] += ";"
		}

		if i != 0 {
			// remove the insert part and replace it with nothing
			lines[i] = strings.Replace(lines[i], match[0], "", 1)
			// remove the starting newline char that we got at the begining by spliting at ;
			lines[i] = strings.Replace(lines[i], "\n", "", 1)
			// return newline char and add spaces for the len of insert part, to pretify the output
			lines[i] = "\n" + strings.Repeat(" ", len(match[0])) + lines[i]
		}
	}

	return strings.Join(lines, "")
}

func main() {
	input := `INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG121', 'The Dinner Game', 141, DEFAULT, 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG122', 'The Dinner Game', 142, DEFAULT, 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG123', 'The Dinner Game', 142, DEFAULT, 'Comedy');`
	output := convertSingleLineToMultilineSQL(input)

	fmt.Println(output)
}
