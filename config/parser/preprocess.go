package configparser

import (
	"strings"
	"unicode"
)

func cleanRow(s string) string {
	var b strings.Builder

	for _, r := range s {
		if r == '#' {
			break
		}

		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '=' {

			b.WriteRune(r)
		}
	}

	return b.String()
}

const minRowLen = 3 // because 'K=V'; smaller row make no sense and we drop them

func splitAndCleanRows(s string) []string {
	rows := strings.Split(s, "\n")

	output := make([]string, 0, len(rows))
	for _, row := range rows {
		row = cleanRow(row)
		if len([]rune(row)) < minRowLen {
			continue
		}

		output = append(output, row)
	}

	return output
}
