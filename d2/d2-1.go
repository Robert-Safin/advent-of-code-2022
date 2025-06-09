package d2

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strings"
)

func Solve1() {
	data, err := utils.GetInput("./d2/d2.txt")
	if err != nil {
		log.Fatal(err)
	}
	points := map[string]int{
		"AX": 4,
		"AY": 8,
		"AZ": 3,

		"BX": 1,
		"BY": 5,
		"BZ": 9,

		"CX": 7,
		"CY": 2,
		"CZ": 6,
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
