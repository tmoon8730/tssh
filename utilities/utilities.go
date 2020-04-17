package utilities

// Check simple error panic function
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
