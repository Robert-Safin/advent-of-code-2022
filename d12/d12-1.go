package d12

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strings"
)

type State struct {
	Pos  [2]int
	Step int
}

func Solve1() {
	data, err := utils.GetInput("./d12/d12.txt")
	if err != nil {
		log.Fatal(err)
	}

	game_map, start, end := parse_map(data)

	queue := []State{{Pos: start, Step: 0}}
	visited := map[[2]int]bool{start: true}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.Pos == end {
			fmt.Println("end found at step", state.Step)
			return
		}

		visited[[2]int{state.Pos[0], state.Pos[1]}] = true

		directions := [][2]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}

		for _, dir := range directions {
			new_y := state.Pos[0] + dir[0]
			new_x := state.Pos[1] + dir[1]

			if new_y < 0 || new_x < 0 || new_y >= len(game_map) || new_x >= len(game_map[0]) {
				continue
			}

			newPos := [2]int{new_y, new_x}
			if visited[newPos] {
				continue
			}

			if game_map[new_y][new_x] <= game_map[state.Pos[0]][state.Pos[1]]+1 {
				visited[newPos] = true
				queue = append(queue, State{
					Pos:  newPos,
					Step: state.Step + 1,
				})
			}
		}

	}

}

func parse_map(input string) ([][]int, [2]int, [2]int) {
	lines := strings.Split(input, "\n")
	var matrix [][]int
	var start, end [2]int

	for i := range lines {
		var row []int
		for j, ch := range strings.Split(lines[i], "") {
			switch ch {
			case "S":
				start = [2]int{i, j}
				row = append(row, 1)

			case "E":
				end = [2]int{i, j}
				row = append(row, 26)

			default:
				row = append(row, int(ch[0]-96))
			}
		}
		matrix = append(matrix, row)
	}

	return matrix, start, end
}
