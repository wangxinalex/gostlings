// Concept: width and precision with fmt
// Task: complete padName and formatPrice
// Expected output: [Ada       ] $3.50
// Hint: one verb pads to a fixed width, left-aligned so the spaces follow the
//       name, and the other fixes the number of decimals (Go doc: fmt)
// Stuck?: %-10s for the padded name and %.2f for the price.

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
