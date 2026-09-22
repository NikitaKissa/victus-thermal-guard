package configparser

import (
	"fmt"
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

func parseRow(row string) (key string, value string, err error) {
	elements := strings.Split(row, "=")
	if len(elements) > 2 {
		return "", "", fmt.Errorf(
			"row `%s` has more then 1 `=`, can't parse it to key=value structure: %w",
			row,
			ErrSyntax,
		)
	}

	if len(elements) < 2 {
		return "", "", fmt.Errorf(
			"row `%s` doesn't match key=value structure: %w",
			row,
			ErrSyntax,
		)
	}

	return elements[0], elements[1], nil
}
