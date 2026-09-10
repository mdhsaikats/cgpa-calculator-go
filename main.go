package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string

		fmt.Print("\nEnter Total Semesters (or type 'exit' to quit): ")
		fmt.Scan(&input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			return
		}

		n, err := strconv.Atoi(input)

		if err != nil || n <= 0 {
			fmt.Println("Invalid input. Please enter a valid positive number.")
			continue
		}

		var totalGradePoints float64
		var totalCreditHours float64
		var sgpa, credit float64

		for i := 1; i <= n; i++ {
			fmt.Printf("Enter SGPA for semester %d: ", i)
			_, err := fmt.Scan(&sgpa)
			if err != nil || sgpa < 0 {
				fmt.Println("Invalid SGPA. Please start over.")
				var discard string
				fmt.Scanln(&discard)
				break
			}

			fmt.Printf("Enter Credit Hours for semester %d: ", i)
			_, err = fmt.Scan(&credit)
			if err != nil || credit <= 0 {
				fmt.Println("Invalid Credit Hour. Please start over.")
				var discard string
				fmt.Scanln(&discard)
				break
			}

			totalGradePoints += sgpa * credit
			totalCreditHours += credit
		}

		if totalCreditHours > 0 {
			cgpa := totalGradePoints / totalCreditHours
			fmt.Printf("Your CGPA is: %.2f\n", cgpa)
		}
	}
}
