package d4

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d4/d4.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(data, "\n")

	no_overlap := 0

	for _, line := range lines {
		split := strings.Split(line, ",")
		left_split := strings.Split(split[0], "-")
		right_split := strings.Split(split[1], "-")

		l, _ := strconv.Atoi(left_split[0])
		r, _ := strconv.Atoi(left_split[1])

		left := Range{start: l, end: r}

		l, _ = strconv.Atoi(right_split[0])
		r, _ = strconv.Atoi(right_split[1])
		right := Range{start: l, end: r}

		if left.end < right.start || right.end < left.start {
			no_overlap += 1
		}

	}
	fmt.Println(len(lines) - no_overlap)
}
