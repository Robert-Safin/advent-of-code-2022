package d8

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
)

func Solve2() {
	data, err := utils.GetInput("./d8/d8.txt")
	if err != nil {
		log.Fatal(err)
	}

	max_score := 0

	parsed := parse(data)
	width := len(parsed[0])
	height := len(parsed)

	for row_i, row := range parsed {
		for col_i, v := range row {

			if is_edge(row_i, col_i, width, height) {
				continue
			}

			col := GetColumn(parsed, col_i)

			l, r := get_score_x(v, col_i, row)
			t, b := get_score_y(v, row_i, col)
			score := l * r * t * b

			if score > max_score {
				max_score = score
			}
		}
	}
	fmt.Println(max_score)
}

func get_score_y(v int, r int, col []int) (int, int) {
	top_side := col[:r]
	bottom_side := col[r+1:]

	top_seen := 0
	for i := len(top_side) - 1; i >= 0; i-- {
		top_seen += 1
		if top_side[i] >= v {
			break
		}
	}

	bottom_seen := 0
	for i := 0; i < len(bottom_side); i++ {
		bottom_seen += 1
		if bottom_side[i] >= v {
			break
		}
	}

	return top_seen, bottom_seen
}

func get_score_x(v int, c int, row []int) (int, int) {
	left_side := row[:c]
	right_side := row[c+1:]

	left_seen := 0
	for i := len(left_side) - 1; i >= 0; i-- {
		left_seen += 1
		if left_side[i] >= v {
			break
		}
	}

	right_seen := 0
	for i := 0; i < len(right_side); i++ {
		right_seen += 1
		if right_side[i] >= v {
			break
		}
	}

	return left_seen, right_seen
}
