package delete

import (
	"log"
	"strings"
)

// Remove drops Tomatoes and Mango from the basket
func Remove(container []string) []string {
	log.Println("Executing : v1.AddFruits")

	for i, value := range container {
		switch strings.ToLower(value) {
		case "mango":
			container = append(container[:i], container[1+1:]...)
		case "tomatoes":
			container = append(container[:i], container[1+1:]...)
		}
	}
	return container
}
