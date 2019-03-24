package main

import (
	"github.com/dmigwi/golang-modules/printer"
	"github.com/dmigwi/golang-modules/topup"
)

func main() {
	var basket = []string{"Pineapple", "Tomatoes"}

	topup.AddFruits(basket)

	printer.Display(basket)
}
