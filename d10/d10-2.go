package d10

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d10/d10.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(data, "\n")

	register := 1
	cycle := 0
	screen := ""

	drawPixel := func() {
		pixelPos := cycle % 40
		if pixelPos >= register-1 && pixelPos <= register+1 {
			screen += "#"
		} else {
			screen += "."
		}
		cycle++
		if cycle%40 == 0 {
			screen += "\n"
		}
	}

	for _, line := range lines {
		if line == "noop" {
			drawPixel()
		} else {
			split := strings.Split(line, " ")
			val, _ := strconv.Atoi(split[1])

			// takes 2 cycles
			drawPixel()
			drawPixel()
			register += val
		}
	}

	fmt.Println("CRT Output:")
	fmt.Print(screen)
}
