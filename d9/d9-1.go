package d9

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type State struct {
	head [2]int
	tail [2]int
}

func Solve1() {
	data, err := utils.GetInput("./d9/d9.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := State{
		head: [2]int{0, 0},
		tail: [2]int{0, 0},
	}
	unique_positions := map[[2]int][]bool{}
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		split := strings.Split(line, " ")
		dir := split[0]
		dist, _ := strconv.Atoi(split[1])

		for i := 0; i < dist; i++ {
			move_head(dir, &s)
			if chebyshev_distance(&s) > 1 {
				move_tail(&s)
			}
			unique_positions[s.tail] = []bool{}
		}
	}
	fmt.Println(unique_positions)
	fmt.Println(len(unique_positions))

}

func chebyshev_distance(s *State) int {
	dx := int(math.Abs(float64(s.head[0] - s.tail[0])))
	dy := int(math.Abs(float64(s.head[1] - s.tail[1])))
	return max(dx, dy)
}

func move_tail(s *State) {
	dx := s.head[0] - s.tail[0]
	dy := s.head[1] - s.tail[1]

	if abs(dx) > 1 || abs(dy) > 1 {
		// move one step towards head
		if dx != 0 {
			s.tail[0] += sign(dx)
		}
		if dy != 0 {
			s.tail[1] += sign(dy)
		}
	}
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func move_head(dir string, s *State) {
	if dir == "R" {
		s.head[1] += 1
	}
	if dir == "L" {
		s.head[1] -= 1
	}
	if dir == "U" {
		s.head[0] += 1
	}
	if dir == "D" {
		s.head[0] -= 1
	}
}
