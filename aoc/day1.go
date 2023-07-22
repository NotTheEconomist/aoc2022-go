package aoc

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed day1.txt
var input string

func Day1_solve() (string, string) {
	// Part 1
	var maxSnack int
	elves := strings.Split(input, "\n\n")
	snackValues := make([]int, len(elves))
	for i, elf := range elves {
		var snackTotal int
		snacks := strings.Split(elf, "\n")
		for _, snack := range snacks {
			value, err := strconv.Atoi(snack)
			if err != nil {
				value = 0
			}
			snackTotal += value
		}
		if maxSnack < snackTotal {
			maxSnack = snackTotal
		}
		snackValues[i] = snackTotal
	}

	// Part 2
	sort.Sort(sort.IntSlice(snackValues))
	topThree := []int{snackValues[len(snackValues)-3], snackValues[len(snackValues)-2], snackValues[len(snackValues)-1]}
	topThreeTotal := topThree[0] + topThree[1] + topThree[2]

	return strconv.Itoa(maxSnack), strconv.Itoa(topThreeTotal)
}
