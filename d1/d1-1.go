package d1

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve1() {
	data, err := utils.GetInput("./d1/d1.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(data)
	split := strings.Split(str, "\n\n")

	max := 0

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
		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)

}
