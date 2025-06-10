package d6

import (
	"advent-of-code-2022/utils"
	"log"
)

func Solve2() {
	data, err := utils.GetInput("./d6/d6.txt")
	if err != nil {
		log.Fatal(err)
	}

	length := 14

	count := map[string]int{}

	for i := 0; i < length; i++ {
		count[string(data[i])] += 1
	}

	for i := length; i < len(data); i++ {
		oldChar := string(data[i-length])
		count[oldChar] -= 1
		if count[oldChar] == 0 {
			delete(count, oldChar)
		}

		newChar := string(data[i])
		count[newChar] += 1

		if len(count) == length {
			println(i + 1)
			break
		}
	}

}
