package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"

	RED_MAX   = 12
	GREEN_MAX = 13
	BLUE_MAX  = 14
)

func parseRounds(rounds string) map[string]int {
	roundMap := make(map[string]int)
	for _, round := range strings.Split(rounds, ",") {
		cleanRound := strings.Trim(round, " ")
		countStr, colour, ok := strings.Cut(cleanRound, " ")
		if !ok {
			fmt.Println(errors.New("game did not contain ' '"))
			return nil
		}
		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		roundMap[colour] = count
	}
	return roundMap
}

func getMax(game []map[string]int, colour string) int {
	max := 0
	for _, round := range game {
		if round[colour] > max {
			max = round[colour]
		}
	}
	return max
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

	fileMap := make(map[int][]map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		dirtyID, dirtyRounds, ok := strings.Cut(line, ":")
		if !ok {
			fmt.Println(errors.New("line did not contain colon"))
			return
		}

		strID := strings.Split(dirtyID, " ")[1]
		id, err := strconv.Atoi(strID)
		if err != nil {
			fmt.Println(err)
			return
		}

		lineRounds := strings.Split(dirtyRounds, ";")
		gameMap := make([]map[string]int, 0)
		for _, rounds := range lineRounds {
			roundMap := parseRounds(rounds)
			if roundMap == nil {
				return
			}
			gameMap = append(gameMap, roundMap)
		}
		fileMap[id] = gameMap
	}

	powerArr := make([]int, len(fileMap)+1)
	for i, game := range fileMap {
		redMax := getMax(game, RED)
		blueMax := getMax(game, BLUE)
		greenMax := getMax(game, GREEN)

		powerArr[i] = redMax * blueMax * greenMax
	}

	sum := 0
	for _, power := range powerArr {
		sum += power
	}
	fmt.Println(sum)
}
