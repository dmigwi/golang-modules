package topup

// AddFruits adds several fruits
func AddFruits(container []string) []string {
	items := []string{"Mango", "Banana", "Watermelon"}

	container = append(container, items...)
	return container
}
