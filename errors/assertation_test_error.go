package errors

import (
	"fmt"
	"log"
)

// AssertionTestError Handles a fatal error for assertion mistake
// it receives as parameter the test number, the expected output and
// the function returned result, respectively
func (e *Errors) AssertionTestError(testNumber int, expectedOutput interface{}, testedReturnOutput interface{}) {
	o := expectedOutput
	tro := testedReturnOutput

	fmt.Printf("\nAssertation failed on test number %v!!\n", testNumber)
	err := fmt.Errorf("\n Expected -> %v\n Returned -> %v\n", o, tro)
	log.Fatalln(err)
}
