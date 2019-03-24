package topup

import (
	"log"
)

// AddFruits adds several fruits
func AddFruits(container []string) []string {
	log.Println("Executing : github.com/dmigwi/golang-modules/topup.AddFruits")

	items := []string{"Mango", "Banana", "Watermelon"}

	container = append(container, items...)
	return container
}
