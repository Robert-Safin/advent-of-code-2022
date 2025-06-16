package d11

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Operation struct {
	Left     string
	Operator string
	Right    string
}

type Monkey struct {
	Items              []int
	Operation_operator string
	Operation          Operation
	test               int
	True_throw         int
	False_throw        int
}

func Solve1() {
	data, err := utils.GetInput("./d11/d11.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(data, "\n\n")
	monkeys := parse_monekys(lines)

	counts := map[int]int{}

	for i := 0; i < 20; i++ {
		for i := range monkeys {
			monkey := &monkeys[i]

			for _, item := range monkey.Items {
				counts[i] += 1
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

				newLevel = newLevel / 3

				if newLevel%monkey.test == 0 {
					monkeys[monkey.True_throw].Items = append(monkeys[monkey.True_throw].Items, newLevel)
				} else {
					monkeys[monkey.False_throw].Items = append(monkeys[monkey.False_throw].Items, newLevel)
				}
			}
			monkey.Items = []int{}
		}

	}

	count_vals := []int{}
	for _, v := range counts {
		count_vals = append(count_vals, v)
	}

	sort.Slice(count_vals, func(i, j int) bool {
		return count_vals[i] > count_vals[j]
	})

	fmt.Println(count_vals[0] * count_vals[1])

}

func parse_monekys(lines []string) []Monkey {
	var monkeys []Monkey

	for _, m := range lines {
		var monkey Monkey
		split := strings.Split(m, "\n")
		items_split := strings.Split(split[1], ":")
		items_num_str := items_split[1]
		num_strs := strings.Split(items_num_str, ",")

		var items []int
		for _, s := range num_strs {
			striped := strings.ReplaceAll(s, " ", "")
			n, _ := strconv.Atoi(striped)

			items = append(items, n)
		}
		monkey.Items = items

		op_str := split[2]

		op_str_split := strings.Split(op_str, "=")
		op_str_split = strings.Split(op_str_split[1], " ")
		monkey.Operation.Left = op_str_split[1]
		monkey.Operation.Operator = op_str_split[2]
		monkey.Operation.Right = op_str_split[3]

		test := split[3]
		test_split := strings.Split(test, " ")
		test_n, _ := strconv.Atoi(test_split[5])
		monkey.test = test_n

		true_throw := split[4]
		true_split := strings.Split(true_throw, " ")
		true_n, _ := strconv.Atoi(true_split[9])
		monkey.True_throw = true_n

		false_throw := split[5]
		false_split := strings.Split(false_throw, " ")
		false_n, _ := strconv.Atoi(false_split[9])
		monkey.False_throw = false_n

		monkeys = append(monkeys, monkey)

	}

	return monkeys
}
