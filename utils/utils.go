package utils

import (
	"bufio"
	"os"
)

func ReadInput(day string) ([]string, error) {
	f, err := os.Open(day + "/input.txt")
	if err != nil {
        return nil, err
    }
	defer f.Close()

	scan := bufio.NewScanner(f)
	lines := make([]string,0)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
    return lines, nil
}
