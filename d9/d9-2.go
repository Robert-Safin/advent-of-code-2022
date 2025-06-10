package d9

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d9/d9.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize rope with 10 knots at (0, 0)
	rope := make([][2]int, 10)
	visited := map[[2]int]struct{}{}

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		dir := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		for i := 0; i < steps; i++ {
			move(&rope[0], dir) // move head

			// Move each knot toward the previous one
			for j := 1; j < len(rope); j++ {
				move_knot(&rope[j], rope[j-1])
			}

			// Mark tail position as visited
			visited[rope[9]] = struct{}{}
		}
	}

	fmt.Println("Unique tail positions:", len(visited))
}

func move(pos *[2]int, dir string) {
	switch dir {
	case "R":
		pos[1] += 1
	case "L":
		pos[1] -= 1
	case "U":
		pos[0] += 1
	case "D":
		pos[0] -= 1
	}
}

func move_knot(curr *[2]int, prev [2]int) {
	dx := prev[0] - curr[0]
	dy := prev[1] - curr[1]

	if max(abs(dx), abs(dy)) > 1 {
		curr[0] += sign(dx)
		curr[1] += sign(dy)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
