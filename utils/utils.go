package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const url = "https://adventofcode.com/2023"

func GetInput(day string) ([]string, error) {
	lines, err := ReadInput(day)
	if err != nil { // Try to download
		fmt.Println("could not find input.txt. downloading...")
		numeral, _ := strings.CutPrefix(day, "day")
		req, e := http.NewRequest("GET", url+"/day/"+numeral+"/input", nil)
		if e != nil {
			return nil, e
		}
		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: os.Getenv("SESSION"),
		})
		rsp, e := http.DefaultClient.Do(req)
		if e != nil {
			return nil, e
		}
		defer rsp.Body.Close()
		body, e := io.ReadAll(rsp.Body)
		if e != nil {
			return nil, e
		}
		os.WriteFile(day+"/input.txt", body, 0644)
		return strings.Split(string(body), "\n"), nil
	} else {
		return lines, nil
	}
}

func ReadInput(day string) ([]string, error) {
	file := day + "/input.txt"
	if strings.ContainsRune(day, '/') {
		file = day
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines, nil
}

func ReadInputAsString(day string) (string, error) {
	file := day + "/input.txt"
	if strings.ContainsRune(day, '/') {
		file = day
	}
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
