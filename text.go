/*
MIT License

Copyright (c) 2023 A-Boring-Square

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
*/

package lunara_framework

import (
	"fmt"
)

// TextColors struct holds ANSI escape codes for text color formatting.
type TextColors struct {
	White  string
	Red    string
	Blue   string
	Green  string
	Yellow string
}

// InitTextColors initializes and returns a TextColors struct with predefined color codes.
func InitTextColors() TextColors {
	return TextColors{
		White:  "\033[97m",   // White
		Red:    "\033[91m",   // Red
		Blue:   "\033[94m",   // Blue
		Green:  "\033[92m",   // Green
		Yellow: "\033[93m",   // Yellow
	}
}
// PrintColorf prints formatted text with the specified color.
func PrintColorf(text string, color string) {
	fmt.Printf("%s%s\033[0m", color, text) // Use ANSI escape codes to set the color, and reset it after the text.
}

// PrintColorln prints a line of text with the specified color.
func PrintColorln(text string, color string) {
	fmt.Printf("%s%s\033[0m\n", color, text) // Use ANSI escape codes to set the color, and reset it after the text, followed by a newline.
}


/*

Exampel

func main() {
	// Example equations: 2x + 3 = 11 or 11 = 3 + 2x
	coefficients := []float64{2.0, 1.0} // Coefficients of x
	constants := []float64{11.0}        // Constants

	solution, err := SolveTwoStepEquation(coefficients, constants)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Solution: x = %.2f\n", solution)
	}

	// Example with invalid input
	invalidCoefficients := []float64{0.0, 1.0} // Coefficients with a zero value
	invalidConstants := []float64{5.0}         // Constants

	invalidSolution, invalidErr := SolveTwoStepEquation(invalidCoefficients, invalidConstants)

	if invalidErr != nil {
		fmt.Println("Error:", invalidErr)
	}
}

*/