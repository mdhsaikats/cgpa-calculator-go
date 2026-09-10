package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string

		fmt.Println("\nEnter Total Semester (or type 'exit' to quit):")
		fmt.Scan(&input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			return
		}

		n, err := strconv.Atoi(input)

		if err != nil || n <= 0 {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		graderPoint := make([]float64, n)
		creditHour := make([]float64, n)
		var upper float64
		var creditHourSum float64
		var result float64

		for i := 1; i <= n; i++ {
			fmt.Printf("Enter your SGPA for s :- %d ", i)
			fmt.Scan(&graderPoint[i-1])

			fmt.Printf("Enter your Credit Hour for s :- %d ", i)
			fmt.Scan(&creditHour[i-1])

			upper += graderPoint[i-1] * creditHour[i-1]
			creditHourSum += creditHour[i-1]
		}

		if creditHourSum > 0 {
			result = upper / creditHourSum
			fmt.Printf("Your CGPA is: %.2f\n", result)
		} else {
			fmt.Println("Total credit hours cannot be zero.")
		}
	}
}
