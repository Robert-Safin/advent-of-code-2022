package d8

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve1() {
	data, err := utils.GetInput("./d8/d8.txt")
	if err != nil {
		log.Fatal(err)
	}

	visible_count := 0

	parsed := parse(data)
	width := len(parsed[0])
	height := len(parsed)

	for row_i, row := range parsed {
		for col_i, v := range row {

			col := GetColumn(parsed, col_i)

			if is_edge(row_i, col_i, width, height) || check_x_is_visible(v, col_i, row) || check_y_is_visible(v, row_i, col) {
				visible_count += 1
			}
		}
	}

	fmt.Println(visible_count)
}

func GetColumn(matrix [][]int, col int) []int {
	var column []int

	for _, row := range matrix {
		if col < len(row) {
			column = append(column, row[col])
		}
	}
	return column
}

func check_y_is_visible(v int, r int, col []int) bool {
	top_side := col[:r]
	bottom_side := col[r+1:]

	top_blocked := false
	for _, top_neighbour := range top_side {
		if top_neighbour >= v {
			top_blocked = true
		}
	}

	bottom_blocked := false
	for _, bottom_neighbour := range bottom_side {
		if bottom_neighbour >= v {
			bottom_blocked = true
		}
	}

	if bottom_blocked == false || top_blocked == false {
		return true
	}

	return false
}

func check_x_is_visible(v int, c int, row []int) bool {
	left_side := row[:c]
	right_side := row[c+1:]

	left_blocked := false
	for _, left_neighbour := range left_side {
		if left_neighbour >= v {
			left_blocked = true
		}
	}

	right_blocked := false
	for _, right_neighbour := range right_side {
		if right_neighbour >= v {
			right_blocked = true
		}
	}

	if left_blocked == false || right_blocked == false {
		return true
	}

	return false
}

func is_edge(r, c, width, height int) bool {
	if r == 0 || c == 0 || r == height-1 || c == width-1 {
		return true
	}
	return false
}

func parse(data string) [][]int {
	lines := strings.Split(data, "\n")
	parsed := [][]int{}

	for _, row_str := range lines {
		row_int := []int{}
		for _, v := range strings.Split(row_str, "") {
			int, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			row_int = append(row_int, int)
		}
		parsed = append(parsed, row_int)
	}
	return parsed
}
