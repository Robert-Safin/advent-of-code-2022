package d7

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strings"
)

func Solve2() {
	data, err := utils.GetInput("./d7/d7.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(data, "\n")
	lines = lines[1:]

	root := Dir{
		name:     "/",
		parent:   nil,
		children: map[string]*Dir{},
		files:    []File{},
	}
	current_dir := &root

	for i := 0; i < len(lines); i++ {
		if lines[i] == "$ ls" {
			start, end := return_ls_chunk(lines, i)
			parse_ls_chunk(lines, start, end, current_dir)
			continue
		}

		if strings.HasPrefix(lines[i], "$ cd ") {
			parts := strings.Split(lines[i], " ")
			if len(parts) < 3 {
				log.Fatalf("unexpected cd line format: %q", lines[i])
			}
			target := parts[2]
			current_dir = handle_cd(current_dir, target)
		}

	}

	var sizes []int
	TraverseAndSumSizes(&root, &sizes)
	fs_max := 70000000
	space_required := 30000000
	used := sizes[len(sizes)-1]
	free := fs_max - used
	space_to_delete := space_required - free

	min := space_required
	for _, s := range sizes {
		if s >= space_to_delete && s < min {
			min = s
		}
	}

	fmt.Println(min)

}
