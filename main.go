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

	r := regexp.MustCompile(`(?i)insert into \w+(.*) values `)
	match := r.FindAllString(lines[0], -1)

	for i := range len(lines) {
		if i != len(lines)-1 {
			lines[i] += ","
		} else {
			lines[i] += ";"
		}

		if i != 0 {
			lines[i] = strings.Replace(lines[i], match[0], "", 1)
		}
	}

	return strings.Join(lines, "")
}

func main() {
	input := `INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`
	output := convertSingleLineToMultilineSQL(input)

	fmt.Println(output)
}
