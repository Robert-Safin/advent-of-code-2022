package d2

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d2/d2.txt")
	if err != nil {
		log.Fatal(err)
	}
	points := map[string]int{
		"AX": 0 + 3,
		"AY": 3 + 1,
		"AZ": 6 + 2,

		"BX": 0 + 1,
		"BY": 3 + 2,
		"BZ": 6 + 3,

		"CX": 0 + 2,
		"CY": 3 + 3,
		"CZ": 6 + 1,
	}
	games := strings.Split(data, "\n")
	total := 0
	for _, g := range games {
		hands := strings.Split(g, " ")
		key := hands[0] + hands[1]
		score, _ := points[key]
		total += score
	}
	fmt.Println(total)
}
