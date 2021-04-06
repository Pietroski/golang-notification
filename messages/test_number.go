package messages

import "fmt"

// TestNumberExecutionTime returns a formatted string of the current
// test number passed as parameter
func TestNumberExecutionTime(testNumber int) string {
	ps := fmt.Sprintf("Test number %v took", testNumber)
	return ps
}
