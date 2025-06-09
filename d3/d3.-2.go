package d3

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d3/d3.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(data, "\n")
	groups := [][]string{}

	sub := []string{}
	for i, _ := range lines {
		sub = append(sub, lines[i])
		if len(sub) == 3 {
			groups = append(groups, sub)
			sub = []string{}
		}
	}

	score := 0

	for _, group := range groups {
		intersection := []rune{}

		for _, c := range group[0] {
			char := string(c)
			if strings.Contains(group[1], char) && strings.Contains(group[2], char) {
				intersection = append(intersection, []rune(char)[0])
			}
		}
		r := intersection[0]
		if r >= 'a' && r <= 'z' {
			score += int(r - 'a' + 1)
		}
		if r >= 'A' && r <= 'Z' {
			score += int(r - 'A' + 27)
		}
	}
	fmt.Println(score)
}
