package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)                // fire prompt for input
	input, err := r.ReadString('\n') // read terminal input

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create new bill name: ", reader)

	billObject := newBill(name)
	fmt.Println("Created the bill - ", billObject.name)

	return billObject
}

func promptOption(billObject bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, t - add tip, s - save bill: ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		price2float, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOption(billObject)
		}

		billObject.addItem(name, price2float) // update bill items
		fmt.Println(name, price, ":..Item added!")
		promptOption(billObject)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		tip2float, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOption(billObject)
		}

		billObject.updateTip(tip2float) // update bill tip
		fmt.Println("You gave a tip of: ", tip)
		promptOption(billObject)

	case "s":
		billObject.save() // save file
		fmt.Println("You saved the bill - ", billObject.name)

	default:
		fmt.Println("That was not a valid option...")
		promptOption(billObject)
	}

}

func main() {
	myBill := createBill()
	promptOption(myBill)
}
