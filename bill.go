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

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format bill
func (b *bill) format() string {
	var total float64
	format_string := "Bill Breakdown: \n"

	// list items
	for key, value := range b.items {
		format_string += fmt.Sprintf("%-25v ...%0.2f \n", key+":", value)
		total += value
	}

	// tips
	format_string += fmt.Sprintf("%-25v ...%0.2f \n", "tip:", b.tip)

	//total
	format_string += fmt.Sprintf("%-25v ...%0.2f", "total:", total+b.tip)

	return format_string
}

// Update tip
func (b *bill) updateTip(value float64) {
	b.tip = value
	// (*b).tip = value
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
