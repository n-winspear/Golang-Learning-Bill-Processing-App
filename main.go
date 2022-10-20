package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "Create a new bill name: ")

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput(reader, "Choose option (a - add item, s - save bill, t - add tip): ")

	switch opt {
	case "a":
		name, _ := getInput(reader, "Item name: ")
		price, _ := getInput(reader, "Item price: ")

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Price must be a number.")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println(b.formatBill())
	case "t":
		tip, _ := getInput(reader, "Enter tip amount (£): ")

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Tip must be a number.")
			promptOptions(b)
		}

		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promptOptions(b)
	default:
		fmt.Println("Not a valid option...")
		promptOptions(b)
	}
}

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func main() {
	myBill := createBill()
	promptOptions(myBill)

}
