package printer

import (
	"fmt"
	"log"
)

// Display prints the contents of the container.
func Display(container []string) {
	log.Println("Executing : github.com/dmigwi/golang-modules/printer.Display")

	for i, name := range container {
		fmt.Printf("%s is at index %d \n", name, i)
	}
}
