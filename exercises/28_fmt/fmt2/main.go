// Concept: width and precision with fmt
// Task: complete padName and formatPrice
// Expected output: [Ada       ] $3.50
// Hint: fmt.Sprintf("%-10s", name) left-aligns to width 10; fmt.Sprintf("%.2f", price) (Go doc: fmt)

package main

import "fmt"

func padName(name string) string {
	// TODO: Return name left-aligned to width 10.
	return name
}

func formatPrice(price float64) string {
	// TODO: Return price with two decimal places.
	return fmt.Sprint(price)
}

func main() {
	fmt.Printf("[%s] $%s\n", padName("Ada"), formatPrice(3.5))
}
