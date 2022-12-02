package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Shape int
type Goal int

const (
	Rock     = Shape(1)
	Paper    = Shape(2)
	Scissors = Shape(3)
	Win      = Goal(6)
	Lose     = Goal(0)
	Draw     = Goal(3)
)

func main() {
	// Read the input file
	file_data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	// Split the file by newlines
	file_data_split := strings.Split(string(file_data), "\n")

	// Part 1
	rpsMap1 := map[string]Shape{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
	score1 := 0
	// Loop through the file data
	for i := 0; i < len(file_data_split); i++ {
		if file_data_split[i] == "" {
			continue
		}
		// Split the line by spaces
		line := strings.Split(file_data_split[i], " ")
		// Get the first and second shape
		opponentShape := rpsMap1[line[0]]
		playerShape := rpsMap1[line[1]]
		score1 += int(playerShape)
		if playerShape == opponentShape {
			score1 += 3
		}
		if playerShape == Rock && opponentShape == Scissors {
			score1 += 6
		}
		if playerShape == Paper && opponentShape == Rock {
			score1 += 6
		}
		if playerShape == Scissors && opponentShape == Paper {
			score1 += 6
		}
	}
	fmt.Println("Part 1 score:", score1)

	// Part 2
	rpsMap2a := map[string]Shape{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	rpsMap2b := map[string]Goal{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}
	score2 := 0
	// Loop through the file data
	for i := 0; i < len(file_data_split); i++ {
		if file_data_split[i] == "" {
			continue
		}
		// Split the line by spaces
		line := strings.Split(file_data_split[i], " ")
		// Get the first and second shape
		opponentShape := rpsMap2a[line[0]]
		playerGoal := rpsMap2b[line[1]]
		playerShape := Shape(0)
		if playerGoal == Win {
			if opponentShape == Rock {
				playerShape = Paper
			}
			if opponentShape == Paper {
				playerShape = Scissors
			}
			if opponentShape == Scissors {
				playerShape = Rock
			}
		}
		if playerGoal == Lose {
			if opponentShape == Rock {
				playerShape = Scissors
			}
			if opponentShape == Paper {
				playerShape = Rock
			}
			if opponentShape == Scissors {
				playerShape = Paper
			}
		}
		if playerGoal == Draw {
			playerShape = opponentShape
		}
		score2 += int(playerShape)
		if playerShape == opponentShape {
			score2 += 3
		}
		if playerShape == Rock && opponentShape == Scissors {
			score2 += 6
		}
		if playerShape == Paper && opponentShape == Rock {
			score2 += 6
		}
		if playerShape == Scissors && opponentShape == Paper {
			score2 += 6
		}
	}
	fmt.Println("Part 2 score:", score2)
}
