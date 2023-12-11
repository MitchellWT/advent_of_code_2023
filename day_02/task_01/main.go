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

func impossibleGames(games map[int][]map[string]int) []int {
	impossible := make([]int, 0)
	for id, game := range games {
		for _, round := range game {
			if round[RED] > RED_MAX || round[BLUE] > BLUE_MAX || round[GREEN] > GREEN_MAX {
				impossible = append(impossible, id)
				break
			}
		}
	}
	return impossible
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
	impossibleIDs := impossibleGames(fileMap)
	sum := 0
	for id := range fileMap {
		isImpossible := false
		for _, impossibleID := range impossibleIDs {
			if id == impossibleID {
				isImpossible = true
				break
			}
		}
		if !isImpossible {
			sum += id
		}
	}
	fmt.Println(sum)
}
