package d5

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d5/d5.txt")
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(data, "\n\n")
	lines := strings.Split(split[0], "\n")
	lines = lines[:len(lines)-1]
	colLine := lines[len(lines)-1]

	colParts := strings.Fields(colLine)
	numCols := len(colParts)

	stacks := make([][]string, numCols)

	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	for _, line := range lines {
		for col := 0; col < numCols; col++ {
			idx := 1 + col*4
			if idx < len(line) && line[idx] != ' ' {
				stacks[col] = append(stacks[col], string(line[idx]))
			}
		}
	}

	moves := strings.Split(split[1], "\n")

	for _, move := range moves {
		split := strings.Split(move, " ")
		move, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		m := Move{
			move: move,
			from: from,
			to:   to,
		}

		fromStack := stacks[from-1]

		moving := fromStack[len(fromStack)-m.move:]

		stacks[m.from-1] = fromStack[:len(fromStack)-m.move]

		stacks[m.to-1] = append(stacks[m.to-1], moving...)

	}

	res := ""

	for _, stack := range stacks {
		res += stack[len(stack)-1]
	}
	fmt.Println(res)
}
