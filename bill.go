package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {

	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

func (b bill) formatBill() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v ...£%v \n", k+":", v)

		total += v
	}

	fs += fmt.Sprintf("%-15v ...£%0.2f", "total: ", total)

	return fs
}
