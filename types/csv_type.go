package types

import (
	"strings"
	"strconv"
)

type CsvInt string

func (s CsvInt) Parse() []int {
	var result []int
	for _, v := range strings.Split(string(s), ",") {
		i, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			continue
		}
		result = append(result, i)
	}

	return result
}

type CsvString string

func (s CsvString) Parse() []string {
	var result []string
	for _, v := range strings.Split(string(s), ",") {
		result = append(result, strings.TrimSpace(v))
	}

	return result
}