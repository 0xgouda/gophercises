package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var score = 0
func startQuiz(data [][]string, pipe chan bool) {
	for _, record := range data {
		fmt.Printf("%s: ", record[0])

		var answer string
		_, err := fmt.Scanf("%s\n", &answer)
		checkErr(err)

		if answer == record[1] {
			fmt.Println("Correct!")
			score += 1
		} else {
			fmt.Println("Incorrect Answer.")
		}
	}
	pipe <- true 
}

func main() {
	fileName := flag.String("file", "problems.csv", "problems csv file")
	timeLimit := flag.Int("limit", 30, "quiz time limit")
	flag.Parse()

	file, err := os.Open(*fileName)
	checkErr(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	checkErr(err)

	cnt := len(data)	
	if cnt == 0 {
		fmt.Println("No questions in file")
		return
	}

	pipe := make(chan bool, 1)
	go startQuiz(data, pipe)

	select {
	case <-pipe:
		fmt.Printf("you scored %d out of %d\n", score, cnt)
	case <-time.After(time.Duration(*timeLimit) * time.Second):
		fmt.Printf("\nTime Limit!!!: you scored %d out of %d\n", score, cnt)
	}
}