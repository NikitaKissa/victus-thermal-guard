package temperatures

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bytesToInt(asciiTemp []byte) (int, error) {
	str := string(asciiTemp)
	str = strings.TrimSuffix(str, "\n")

	integer, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf(
			"failed translation of []byte{%v} to int: %w \nstringrepresentation of that data {%s}",
			asciiTemp,
			err,
			str,
		)
	}

	return integer, nil
}

func millidegToDeg(i int) float64 {
	return float64(i) / 1000
}

func getTemperature(filepath string) (float64, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return 0, fmt.Errorf(
			"failed to read temperature: %w",
			err,
		)
	}

	millidegrees, err := bytesToInt(data)
	if err != nil {

	}

	return millidegToDeg(millidegrees), err
}
