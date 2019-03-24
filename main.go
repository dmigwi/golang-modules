package main

import (
	"github.com/dmigwi/golang-modules/printer"
	"github.com/dmigwi/golang-modules/topup"
	"github.com/dmigwi/golang-modules/v1/delete"
)

func main() {
	var basket = []string{"Pineapple", "Tomatoes"}

	basket = topup.AddFruits(basket)

	printer.Display(basket)

	basket = delete.Remove(basket)

	printer.Display(basket)
}
