package d3

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strings"
)

func Solve1() {
	data, err := utils.GetInput("./d3/d3.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(data, "\n")

	score := 0

	for _, sack := range lines {
		mid := len(sack) / 2
		left := sack[:mid]
		right := sack[mid:]

		left_set := map[rune]bool{}
		for _, ch := range left {
			left_set[ch] = true
		}
		right_set := map[rune]bool{}
		for _, ch := range right {
			right_set[ch] = true
		}

		intersection := []rune{}
		for kl, _ := range left_set {
			for kr, _ := range right_set {
				if kl == kr {
					intersection = append(intersection, kr)
				}
			}
		}

		for _, r := range intersection {
			if r >= 'a' && r <= 'z' {
				score += int(r - 'a' + 1)
			}
			if r >= 'A' && r <= 'Z' {
				score += int(r - 'A' + 27)
			}
		}

	}
	fmt.Println(score)
}
