package errors

import (
	"fmt"
	"github.com/Pietroski/notification/errors/models"
	"log"
)

// AssertionTestError Handles a fatal error for assertion mistake
// it receives as parameter the test number, the expected output and
// the function returned result, respectively
func (e *Errors) AssertionTestError(testNumber int, outputs models.VariadicParametersModel) {
	o := outputs.Output
	fro := outputs.FunctionReturnOutput

	fmt.Printf("\nAssertation failed on test number %v!!\n", testNumber)
	err := fmt.Errorf("\n Expected -> %v\n Returned -> %v\n", o, fro)
	log.Fatalln(err)
}
