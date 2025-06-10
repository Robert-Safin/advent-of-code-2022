package d7

import (
	"advent-of-code-2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name     string
	parent   *Dir
	children map[string]*Dir
	files    []File
}

func Solve1() {
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
	sum := 0
	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}

	fmt.Println(sum)

}

func return_ls_chunk(lines []string, current_i int) (int, int) {
	start_i := current_i + 1
	end_i := len(lines) - 1
	for i := current_i + 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "$") {
			end_i = i - 1
			break
		}
	}
	return start_i, end_i
}

func parse_ls_chunk(lines []string, start int, end int, current_dir *Dir) {

	for i := start; i <= end; i++ {
		if strings.Contains(lines[i], "dir") {
			name := strings.Split(lines[i], " ")[1]
			new_dir := Dir{
				name:     name,
				parent:   current_dir,
				children: map[string]*Dir{},
				files:    []File{},
			}
			current_dir.children[name] = &new_dir
		} else {
			split := strings.Split(lines[i], " ")
			size_str := split[0]
			size_int, _ := strconv.Atoi(size_str)

			name := split[1]
			new_file := File{
				size: size_int,
				name: name,
			}
			current_dir.files = append(current_dir.files, new_file)
		}
	}
}

func handle_cd(current_dir *Dir, target string) *Dir {

	if target == ".." {
		return current_dir.parent
	}

	new, present := current_dir.children[target]
	if !present {
		log.Fatal("failed to CD as no target child found")
	}

	return new
}

func TraverseAndSumSizes(dir *Dir, sizes *[]int) int {
	totalSize := 0

	for _, file := range dir.files {
		totalSize += file.size
	}

	for _, child := range dir.children {
		childSize := TraverseAndSumSizes(child, sizes)
		totalSize += childSize
	}

	*sizes = append(*sizes, totalSize)
	return totalSize
}
