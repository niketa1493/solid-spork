package internal

import "fmt"

func PrintResult(result float64) string {
	return "The result is: " + formatFloat(result)
}

func formatFloat(value float64) string {
	return fmt.Sprintf("%.2f", value)
}
