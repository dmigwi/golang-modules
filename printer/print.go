package printer

import "fmt"

// Display prints the contents of the container.
func Display(container []string) {
	for i, name := range container {
		fmt.Printf("%s is at index %d", name, i)
	}
}
