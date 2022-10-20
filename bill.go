package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {

	b := bill{
		name:  name,
		items: map[string]float64{},
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

	total += b.tip
	fs += fmt.Sprintf("%-15v ...£%0.2f\n", "tip: ", b.tip)

	fs += fmt.Sprintf("%-15v ...£%0.2f\n", "total: ", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.formatBill())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Bill {%v} saved to file.\n", b.name)
}
