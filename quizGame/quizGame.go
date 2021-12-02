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
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open %s \n", *csvFile))
	}
	defer file.Close()

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil{
		exit("Unable to read csv file")
	}

	problems := parseLines(csvLines)
	score := quiz(problems)
	fmt.Printf("You score %v out of %v\n", score, len(csvLines))

} 

func quiz(problems []problem) int {
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