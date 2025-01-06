package main

import (
	"fmt"
	"time"
)

type Question struct {
	question string
	options  [4]string
	answer   int
}

func main() {
	// Define a slice of structs to store questions
	questions := []Question{
		{
			question: "What is the capital of France?",
			options:  [4]string{"1. Berlin", "2. Madrid", "3. Paris", "4. Rome"},
			answer:   3,
		},
		{
			question: "Which planet is known as the Red Planet?",
			options:  [4]string{"1. Earth", "2. Mars", "3. Jupiter", "4. Saturn"},
			answer:   2,
		},
		{
			question: "What is the largest ocean on Earth?",
			options:  [4]string{"1. Atlantic", "2. Pacific", "3. Indian", "4. Arctic"},
			answer:   2,
		},
	}

	var score int
	var userAnswer int

	// Start the quiz
	for i, question := range questions {
		// Set a timer for each question
		timer := time.NewTimer(10 * time.Second) // 10 seconds per question
		answerCh := make(chan int)

		// Display question and options
		fmt.Printf("Question %d: %s\n", i+1, question.question)
		for _, option := range question.options {
			fmt.Println(option)
		}

		// Prompt the user for an answer
		go func() {
			var input string
			_, err := fmt.Scan(&input)

			// Handle invalid input
			if err != nil {
				fmt.Println("Invalid input, please enter a valid option number.")
				answerCh <- -1
				return
			}

			// Convert input to integer
			fmt.Sscanf(input, "%d", &userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case userAnswer = <-answerCh:
			if userAnswer == -1 {
				continue
			}
			if userAnswer == question.answer {
				score++
			}
		case <-timer.C:
			fmt.Println("Time's up for this question!")
		}

		// Ask for exit command
		var exit string
		fmt.Print("Type 'exit' to quit or press Enter to continue: ")
		fmt.Scanln(&exit)
		if exit == "exit" {
			fmt.Println("Exiting quiz...")
			break
		}
	}

	// Calculate the score and classify performance
	fmt.Printf("Your final score is %d/%d\n", score, len(questions))
	switch {
	case score == len(questions):
		fmt.Println("Excellent!")
	case score >= len(questions)/2:
		fmt.Println("Good!")
	default:
		fmt.Println("Needs Improvement.")
	}
}
