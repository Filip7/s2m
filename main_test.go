package main

import (
	"testing"
)

var input = `INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
	INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`

var output = `INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
	('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`

func TestMain(t *testing.T) {
	out := convertSingleLineToMultilineSQL(input)

	if out != output {
		t.Errorf("Expected\n\"%s\"\ngot\n\"%s\"", output, out)
	}
}

// func Test_convertSingleLineToMultilineSQL(t *testing.T) {
// 	tests := []struct {
// 		name string // description of this test case
// 		// Named input parameters for target function.
// 		input string
// 		want  string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := convertSingleLineToMultilineSQL(tt.input)
// 			// TODO: update the condition below to compare got with tt.want.
// 			if true {
// 				t.Errorf("convertSingleLineToMultilineSQL() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
