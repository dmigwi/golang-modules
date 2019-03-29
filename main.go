package main

import (
	"yes/v1/delete"

	"github.com/dmigwi/golang-modules/printer"
	"github.com/dmigwi/golang-modules/topup"
)

func main() {
	var basket = []string{"Pineapple", "Tomatoes"}

	basket = topup.AddFruits(basket)

	printer.Display(basket)

	basket = delete.Remove(basket)

	printer.Display(basket)
}
