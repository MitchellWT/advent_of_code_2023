package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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
		for _, char := range scanner.Text() {
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
