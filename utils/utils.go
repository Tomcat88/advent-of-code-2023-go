package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const url = "https://adventofcode.com/2023"

func GetInput(day string) ([]string, error) {
	content, err := GetInputAsString(day)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0)
	for _, l := range strings.Split(content, "\n") {
		if len(l) == 0 {
			continue
		}

		ret = append(ret, l)
	}
	return ret, nil
}

func GetInputAsString(day string) (string, error) {
	lines, err := ReadInputAsString(day)
	if err != nil { // Try to download
		fmt.Println("could not find input.txt. downloading...")
		numeral, _ := strings.CutPrefix(day, "day")
		url := url + "/day/" + numeral + "/input"
		fmt.Println("GET", url)
		req, e := http.NewRequest("GET", url, nil)
		if e != nil {
			return "", e
		}
		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: os.Getenv("TOKEN"),
		})
		rsp, e := http.DefaultClient.Do(req)
		if e != nil {
			return "", e
		}
		defer rsp.Body.Close()
		body, e := io.ReadAll(rsp.Body)
		if e != nil {
			return "", e
		}
		os.WriteFile(day+"/input.txt", body, 0644)
		return string(body), nil
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

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	gcd := Gcd(a, b)
	return (a * b) / gcd
}

func LcmSlice(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	lcm := numbers[0]
	for i := 1; i < len(numbers); i++ {
		lcm = Lcm(lcm, numbers[i])
	}

	return lcm
}

func AtoiSplit(str string, sep string) ([]int, error) {
	slice := strings.Split(str, sep)
	s := make([]int, len(slice))
	for i, n := range slice {
		if n == "" {
			continue
		}
		int, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		s[i] = int
	}
	return s, nil
}
