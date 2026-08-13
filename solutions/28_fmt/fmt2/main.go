// Concept: width and precision with fmt
// Task: complete padName and formatPrice
// Expected output: [Ada       ] $3.50
// Hint: fmt.Sprintf("%-10s", name) left-aligns to width 10; fmt.Sprintf("%.2f", price) (Go doc: fmt)

package main

import "fmt"

func padName(name string) string {
	return fmt.Sprintf("%-10s", name)
}

func formatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}

func main() {
	fmt.Printf("[%s] $%s\n", padName("Ada"), formatPrice(3.5))
}
