package temperatures

import (
	"os"
	"strconv"
	"strings"
)

const (
	NaN         = -1
	ErrReadFile = -2
)

func bytesToInt(asciiTemp []byte) int {
	str := string(asciiTemp)
	str = strings.TrimSuffix(str, "\n")
	integer, err := strconv.Atoi(str)
	if err != nil {
		integer = NaN
	}

	return integer
}

func millidegToDeg(i int) int {
	return i / 1000
}

func getTemperature(filepath string) int {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return ErrReadFile
	}

	millidegrees := bytesToInt(data)
	return millidegToDeg(millidegrees)
}
