package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func recordsFromCSV(filePath *string) [][]string {
	fmt.Println("Loading CSV file: " + *filePath)
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+*filePath, err)
	}
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+*filePath, err)
	}
	return records
}

func getNextQuestion(records [][]string, i int) string {
	return records[i][0]
}

func parseUserInputs() (*string, *int) {
	// parses user input specified in the terminal
	var filePath = flag.String("file", "./problems.csv", "Path to quiz CSV file")
	var timer = flag.Int("timer", 30, "Maximum timer for the quiz")
	flag.Parse()
	return filePath, timer
}
func main() {
	filePath, timer_lim := parseUserInputs()
	fmt.Printf("Game starting. You have %d seconds to finish. Good Luck!", *timer_lim)
	records := recordsFromCSV(filePath)
	totalCorrect := 0
	totalNumberOfQuestions := len(records)
	timer := time.NewTimer(time.Second * time.Duration(*timer_lim))
	defer timer.Stop()

	for i := range records {
		fmt.Println(getNextQuestion(records, i))
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("You got " + strconv.Itoa(totalCorrect) + " answers correct from " + strconv.Itoa(totalNumberOfQuestions) + " questions!")
			return
		case answer := <-answerCh:
			if answer == records[i][1] {
				totalCorrect++
			}
		}

	}

	fmt.Println("You got " + strconv.Itoa(totalCorrect) + " answers correct from " + strconv.Itoa(totalNumberOfQuestions) + " questions!")
}
