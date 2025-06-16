package d11

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d11/d11.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(data, "\n\n")
	monkeys := parse_monekys(lines)

	counts := map[int]int{}

	// 🧠 Calculate modulo = product of all monkey.test values
	modulo := 1
	for _, m := range monkeys {
		modulo *= m.test
	}

	for round := 0; round < 10000; round++ {
		for i := range monkeys {
			monkey := &monkeys[i]

			for _, item := range monkey.Items {
				counts[i]++

				var newLevel int
				if monkey.Operation.Left == "old" && monkey.Operation.Right == "old" {
					newLevel = item * item
				} else {
					right, err := strconv.Atoi(monkey.Operation.Right)
					if err != nil {
						log.Fatal(err)
					}
					switch monkey.Operation.Operator {
					case "*":
						newLevel = item * right
					case "+":
						newLevel = item + right
					case "-":
						newLevel = item - right
					}
				}

				// ✅ Modulo to keep numbers manageable
				newLevel = newLevel % modulo

				if newLevel%monkey.test == 0 {
					monkeys[monkey.True_throw].Items = append(monkeys[monkey.True_throw].Items, newLevel)
				} else {
					monkeys[monkey.False_throw].Items = append(monkeys[monkey.False_throw].Items, newLevel)
				}
			}

			// 🧹 Clear items after processing
			monkey.Items = []int{}
		}
	}

	// 🔢 Extract and sort inspection counts
	countVals := []int{}
	for _, v := range counts {
		countVals = append(countVals, v)
	}

	sort.Slice(countVals, func(i, j int) bool {
		return countVals[i] > countVals[j]
	})

	// 🧾 Output monkey business level
	fmt.Println(countVals[0] * countVals[1])
}
