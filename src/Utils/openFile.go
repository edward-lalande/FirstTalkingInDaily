package Utils

import (
	"os"
	"strings"
)

func OpenFile(filePath string, sep string) ([]string, error) {
	b, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	var content string = string(b)

	return strings.Split(content, sep), nil
}
