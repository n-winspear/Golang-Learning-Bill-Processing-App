package main

import "fmt"

func main() {

	myBill := newBill("Nathan's Bill")

	fmt.Println(myBill.formatBill())

}
