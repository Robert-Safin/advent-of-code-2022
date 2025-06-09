package d1

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d1/d1.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(data)
	split := strings.Split(str, "\n\n")

	maxes := []int{}

	for _, s := range split {
		sum := 0
		for _, v := range strings.Split(s, "\n") {
			if v == "" {
				continue
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
		}
		maxes = append(maxes, sum)
	}
	sort.Slice(maxes, func(i, j int) bool {
		if maxes[i] > maxes[j] {
			return true
		}
		return false
	})

	top_3 := maxes[0] + maxes[1] + maxes[2]
	fmt.Println(top_3)
}
