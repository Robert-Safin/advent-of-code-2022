package d10

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve1() {
	data, err := utils.GetInput("./d10/d10.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(data, "\n")

	register := 1
	cycle := 0
	sumSignalStrengths := 0
	targetCycles := map[int]bool{
		20: true, 60: true, 100: true, 140: true, 180: true, 220: true,
	}

	for _, line := range lines {
		if line == "noop" {
			cycle++
			if targetCycles[cycle] {
				sumSignalStrengths += cycle * register
			}
		} else {
			// addx V takes 2 cycles
			split := strings.Split(line, " ")
			val, _ := strconv.Atoi(split[1])

			for i := 0; i < 2; i++ {
				cycle++
				if targetCycles[cycle] {
					sumSignalStrengths += cycle * register
				}
			}
			register += val // only after 2nd cycle
		}
	}

	fmt.Println("Sum of signal strengths:", sumSignalStrengths)
}
