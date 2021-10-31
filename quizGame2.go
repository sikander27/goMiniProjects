package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)


type problem struct{
	question string
	answer string
}

func main(){
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("timeLimit", 30, "Time limit for quiz")
	flag.Parse()
	// fmt.Printf("sikander \n timelimit %v(%T)", *timeLimit, *timeLimit)
	// Openining the csv file
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open %s \n", *csvFile))
	}
	defer file.Close()  // making sure to close file at the end 

	// Reading from file (all lines)
	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil{
		exit("Unable to read csv file")
	}

	problems := parseLines(csvLines)

	score := quiz(problems) // How can I pass timelimit or the chan to a function
	fmt.Printf("You score %v out of %v\n", score, len(csvLines))

} 

// Main quiz function 
func quiz(problems []problem) int {
	// timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	// <-timer.C
	score := 0
	for i, p := range problems {
		var ans string;
		fmt.Printf("Problem #%v %v = ",(i+1), p.question)
		fmt.Scanf("%s\n", &ans)
	
		if ans == p.answer{
			score ++
		}
	}
	return score
}

// Function to parse line (convert input into custom problem type)
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question : line[0],
			answer : strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}