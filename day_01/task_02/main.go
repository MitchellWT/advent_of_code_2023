package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	ONE   = "one"
	TWO   = "two"
	THREE = "three"
	FOUR  = "four"
	FIVE  = "five"
	SIX   = "six"
	SEVEN = "seven"
	EIGHT = "eight"
	NINE  = "nine"
)

var (
	numStrArr = []string{
		ONE,
		TWO,
		THREE,
		FOUR,
		FIVE,
		SIX,
		SEVEN,
		EIGHT,
		NINE,
	}
	strToNumMap = map[string]int{
		ONE:   1,
		TWO:   2,
		THREE: 3,
		FOUR:  4,
		FIVE:  5,
		SIX:   6,
		SEVEN: 7,
		EIGHT: 8,
		NINE:  9,
	}
)

func startOfNumStr(char rune) bool {
	for _, str := range numStrArr {
		if char == rune(str[0]) {
			return true
		}
	}
	return false
}

func getNumAt(line string, i int) *int {
	str := ""
	for _, numChar := range line[i:] {
		str += string(numChar)
		num, ok := strToNumMap[str]
		if ok {
			return &num
		}
	}
	return nil
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numArr := make([]int, 0)
	for scanner.Scan() {
		lineNumArr := make([]int, 0)
		line := scanner.Text()
		for i, char := range line {
			if startOfNumStr(char) {
				num := getNumAt(line, i)
				if num != nil {
					lineNumArr = append(lineNumArr, *num)
					continue
				}
			}
			if !unicode.IsDigit(char) {
				continue
			}
			i := int(char - '0')
			lineNumArr = append(lineNumArr, i)
		}
		numStr := fmt.Sprintf("%d%d", lineNumArr[0], lineNumArr[len(lineNumArr)-1])
		lineNum, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		numArr = append(numArr, lineNum)
	}
	finalNum := 0
	for _, num := range numArr {
		finalNum += num
	}
	fmt.Println(finalNum)
}
