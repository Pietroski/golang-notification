package errors

var (
	Error ErrorInterface = &Errors{}
)

type ErrorInterface interface {
	CheckPanicError(err error, message ...string)
	CheckFatalError(err error, message ...string)
	AssertionTestError(testNumber int, expectedOutput interface{}, testedReturnOutput interface{})
}

type Errors struct {
	//
}
