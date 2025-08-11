package main

import (
	"testing"
)

func Test_convertSingleLineToMultilineSQL(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input string
		want  string
	}{
		{
			"Simple use case",
			`INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`,
			`INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
                                                             ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`,
		},
		{
			"Different tables use case",
			`INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
INSERT INTO films2 (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films2 (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`,
			`INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
                                                             ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
INSERT INTO films2 (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
                                                              ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`,
		},
		{
			"Two multi, one single in the middle",
			`INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
INSERT INTO test (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films2 (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films2 (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`,
			`INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
                                                             ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
INSERT INTO test (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films2 (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
                                                              ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`,
		},
		{
			"Leave non insert commands",
			`SELECT * FROM films;
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
SELECT * FROM films2;`,
			`SELECT * FROM films;
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
                                                             ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
SELECT * FROM films2;`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertSingleLineToMultilineSQL(tt.input)
			if got != tt.want {
				t.Errorf("Expected\n\"%s\"\ngot\n\"%s\"", tt.want, got)
			}
		})
	}
}
