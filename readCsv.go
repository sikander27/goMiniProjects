package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main(){
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil{
		fmt.Println(err)
	}
	for i, line := range csvLines {
		fmt.Printf("%v -> %v = %v \n", (i+1), line[0], line[1])
	}

}
