package utilities

import "io/ioutil"

// Check simple error panic function
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Helper function to create empty file
func CreateEmptyFile(name string) {
	d := []byte("")
	Check(ioutil.WriteFile(name, d, 0644))
}
