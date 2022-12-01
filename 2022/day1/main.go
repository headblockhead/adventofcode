package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read the input file
	file_data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	// Split the file by newlines
	file_data_split := strings.Split(string(file_data), "\n")
	// Define a list of all the elves
	all_elves := []int{}
	// Define the current elf we are on
	current_elf := 0
	// Loop through the file data
	for i := 0; i < len(file_data_split); i++ {
		// If the line is empty, we know the current elf has been counted
		if file_data_split[i] == "" {
			// If the current elf does not have any calories, we do not add them to the list
			if current_elf != 0 {
				// Add the current elf to the list
				all_elves = append(all_elves, current_elf)
				// Reset the current elf
				current_elf = 0
			}
		} else {
			// Convert the line to an integer
			linecals, err := strconv.Atoi(file_data_split[i])
			if err != nil {
				fmt.Println(err)
			}
			// Add the line to the current elf
			current_elf += linecals
		}
	}
	// Part 1 - Find the highest number of calories
	highest_elf := 0
	for i := 0; i < len(all_elves); i++ {
		if all_elves[i] > highest_elf {
			highest_elf = all_elves[i]
		}
	}
	fmt.Println("highest single elf:", highest_elf)
	// Part 2 - Find the total calories of the top 3 elves
	// Sort the list of elves
	sort.Ints(all_elves)
	// Get the top 3 elves and add their calories together
	fmt.Println("total calories of top 3 elves:", all_elves[len(all_elves)-1]+all_elves[len(all_elves)-2]+all_elves[len(all_elves)-3])
}
