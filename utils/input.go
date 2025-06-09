package utils

import (
	"os"
	"strings"
)

func GetInput(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	str := string(data)
	str = strings.TrimRight(str, "\n")

	return str, nil
}
