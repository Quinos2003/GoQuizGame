package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Question struct holds each quiz question, possible answers, and the correct answer index.
type Question struct {
	question string
	options  []string
	answer   int
}

// Define a few sample questions.
var questions = []Question{
	{
		question: "What is the capital of France?",
		options:  []string{"1. Berlin", "2. Paris", "3. Madrid", "4. Rome"},
		answer:   2,
	},
	{
		question: "What is 2 + 2?",
		options:  []string{"1. 3", "2. 4", "3. 5", "4. 6"},
		answer:   2,
	},
	{
		question: "What is the largest planet in our solar system?",
		options:  []string{"1. Earth", "2. Mars", "3. Jupiter", "4. Saturn"},
		answer:   3,
	},
}

func main() {
	// Seed the random number generator to shuffle questions each time.
	rand.Seed(time.Now().UnixNano())

	// Shuffle the question order for randomness.
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

	// Variables to store the player's score.
	score := 0
	totalQuestions := len(questions)

	// Introduction
	fmt.Println("Welcome to GoQuizGame! Answer the questions by typing the number of your choice.")
	fmt.Println("Let's get started!")

	// Create a scanner to read user input.
	scanner := bufio.NewScanner(os.Stdin)

	// Loop over the shuffled questions.
	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.question)

		// Display answer options.
		for _, option := range q.options {
			fmt.Println(option)
		}

		// Get user input.
		fmt.Print("Your answer: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// Check if the input is a valid option.
		if input == fmt.Sprintf("%d", q.answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Wrong. The correct answer was %d.\n", q.answer)
		}
	}

	// Show the final score.
	fmt.Printf("\nGame Over! Your score: %d/%d\n", score, totalQuestions)
	if score == totalQuestions {
		fmt.Println("Excellent! You got all answers correct!")
	} else if score > totalQuestions/2 {
		fmt.Println("Good job! You scored well.")
	} else {
		fmt.Println("Keep practicing! Better luck next time.")
	}

	// Show the percentage of correct answers.
	percentage := (float32(score) * 100) / float32(totalQuestions)

	fmt.Printf("You answered %.2f%% correctly\n", percentage)
}
