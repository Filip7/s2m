package main

import (
	"testing"
)

var input = `INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`

var output = `INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
    ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');`

func TestMain(t *testing.T) {
	out := convertSingleLineToMultilineSQL(input)

	if out != output {
		t.Errorf("Expected \"%s\" got \"%s\"", output, out)
	}
}
