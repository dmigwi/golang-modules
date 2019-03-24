package main

import (
	"github.com/dmigwi/golang-modules/v1/printer"
	"github.com/dmigwi/golang-modules/v1/topup"
)

func main() {
	var basket = []string{"Pineapple", "Tomatoes"}

	topup.AddFruits(basket)

	printer.Display(basket)
}
