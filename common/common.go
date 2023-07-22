package common

import (
	"bufio"
	"strings"
)

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func Map[I, O any](input []I, fn func(I) O) []O {
	output := make([]O, len(input))
	for i := range input {
		output[i] = fn(input[i])
	}
	return output
}
